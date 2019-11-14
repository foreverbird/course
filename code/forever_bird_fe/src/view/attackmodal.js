import React, { Component } from 'react';
import { Link } from 'react-router-dom';
import {Modal,Button,List,Pagination} from 'antd';
 
import Bird from '../containers/main/bird';

export default class AttackModal extends Component {

    constructor(props){
        super(props);
    }

    render(){
        const { visible, onCancleM,hasBird, birdList, token, attackId } = this.props;
        console.log(this.props);
        if(hasBird){
            return (
                <div>
                    <Modal className="attack-modal" visible={visible}  onCancel={onCancleM} footer={null}  destroyOnClose={true} width={700}>
                        <p>请选择要参与战斗的鸟</p>
                        <div style={{paddingLeft:20}}>
                        <List dataSource={birdList} itemLayout='vertical'
                            grid={{ gutter: 16, column: 2, xs: 1, md : 2,xl : 2, xxl : 2}}
                            renderItem={item => (<List.Item><Bird info={item} isSelf={4}  {...this.props}/></List.Item>)}
                        >
                        </List>
                    </div>
                    </Modal>
                </div>
            );
        }else{
            const state =  {
                token : this.props.token,
                nickName : null
            };
            const birdmarket = {
                pathname : '/birdmarket',
                state : state
            }
            const catchBird = {
                pathname : '/catchbird',
                state : state,
            }
            return (
                <div>
                    <Modal visible={visible} footer={null} destroyOnClose={true} onCancel={onCancleM}>
                        <p>你当前还没有加密鸟代币，可以通过下面两种方式获取</p>
                        <p>(1) 鸟市购买<Link to={birdmarket}>购买</Link></p>
                        <p>(2) 森林捕捉<Link to={catchBird}>捉鸟</Link></p>
                    </Modal>
                </div>
            );
        }
    }
}