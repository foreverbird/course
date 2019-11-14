package data

/**
 * bird基因文件对应的数据
 */
type BirdGene struct {
	Path string
	Name string
	Desc string
}

/**
 * 部位定义
 */
const (
	BIRD_COLOR  = 0
	BIRD_FUR  = 1
	CREST = 2
	EYE  = 3
	FOOT     = 4
	MOUTH  = 5
	TAIL = 6
	WING  = 7
)

/**
 * 部位与关键字映射
 */
var SvgPathMap = map[int]string{
	BIRD_COLOR:"bird_color",
	BIRD_FUR:"bird_fur",
	CREST:"crest",
	EYE:"eye",
	FOOT:"foot",
	MOUTH:"mouth",
	TAIL:"tail",
	WING:"wing",
}