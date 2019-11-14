package logic

import (
	"bird/data"
	"bird/util"
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"sync"
)

const RESOURCE_PATH = "D://project//birds//course//"
const GENES_JSON_PATH = RESOURCE_PATH + "genes.json"
const SVG_SRC_DIR = RESOURCE_PATH + "svg//"
const BIRD_BASE_CONFIG = SVG_SRC_DIR + "bird_base.svg"
const BIRD_GENES_DIR = SVG_SRC_DIR + "genes//"
const SVG_OUTPUT_DIR = "D://project//birds//svg//"

const GENE_NUM = 8

type BirdGenerator struct{}

var instance *BirdGenerator
var once sync.Once

// 存储bird基因配置文件genes.json数据
var geneMap map[string]map[string]data.BirdGene

// 单例，维护一份内存数据
func BirdGeneratorInstance() *BirdGenerator {
	once.Do(func() {
		instance = &BirdGenerator{}
		instance.LoadAllConfig()
	})
	return instance
}

/**
 * 加载genes.json文件
 */
func (*BirdGenerator) LoadAllConfig() {
	geneMap = make(map[string]map[string]data.BirdGene)

	//读取文件
	genes, err := ioutil.ReadFile(GENES_JSON_PATH)
	if err != nil {
		log.Fatal("load bird gene config failed:", err)
	}

	// 解析json字符串
	err = json.Unmarshal(genes, &geneMap)
	if err != nil {
		log.Fatal("load bird gene config failed:", err)
	}
}

/**
 * 生成svg图片
 */
func (*BirdGenerator) GenerateSvg(tokenId int64, genes string) (string) {
	// 根据基因获取配置信息
	indexerPathMap := parsePath(genes)
	// 根据geneMapping组合svg文件
	return assemble(tokenId, indexerPathMap)
}

/**
 ** 根据geneMapping组合svg文件
 ** @param petSvgDto
 ** @return 文件路径
*/
func assemble(tokenId int64, indexerPathMap map[string]string) (string) {
	// 读取bird模板文件：bird_base.svg
	c, _ := ioutil.ReadFile(BIRD_BASE_CONFIG)
	birdBase := string(c)

	// 遍历所有<部位,文件>map
	for indexer, path := range indexerPathMap {
		// indexer对应部位关键字，如wing
		holder := "#" + indexer

		// path为部位图片的相对路径
		svg, _ := ioutil.ReadFile(BIRD_GENES_DIR + "//" + path)

		// 使用部位svg图片数据替换占位符（如："#wing"）
		birdBase = strings.Replace(birdBase, holder, string(svg), -1)
	}

	// 存储图片
	svgPath := SVG_OUTPUT_DIR + strconv.FormatInt(tokenId, 10) + ".svg"
	util.WriteResourceFile(birdBase, svgPath)

	return svgPath
}

func parsePath(genes string) (map[string]string){
	indexerPathMap := make(map[string]string)
	// 去除基因字符串的"0x"
	genes = genes[2:]
	// 遍历基因字符串
	for i := 0; i < GENE_NUM; i++ {
		key := data.SvgPathMap[i]
		code := genes[i * 8: (i + 1) * 8]
		// 计算对应索引[0,15]
		indexer := calcIndexer(code)
		// 通过索引获取图片相对路径
		svgPath := geneMap[key][indexer].Path
		indexerPathMap[key] = svgPath
	}
	return indexerPathMap
}

/**
 * 通过基因code计算索引
 */
func calcIndexer(code string) string {
	// 转为int
	c, _:= strconv.ParseInt(code, 16, 64)
	indexer := int64(0)
	// 每一位数字相加
	for ; c  > 0; {
		indexer += c % 10
		c /= 10
	}

	indexer %= 16
	return strconv.FormatInt(indexer, 10)
}