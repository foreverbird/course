import React, { Component } from 'react';
import { Form, Input, Button,Tooltip,Icon } from 'antd';
 
import {BASE_URL} from '../util/common';
import {Redirect} from 'react-router-dom';
import 'whatwg-fetch';

export default class RegistrationForm extends React.Component {

    constructor(props){
        super(props);
        this.state = {
            redirect : false
        }
    }
  
    handleSubmit = (e) => {
    e.preventDefault();
    this.props.form.validateFieldsAndScroll((err, values) => {
      if (!err) {
        console.log('Received values of form: ', values);
        let Web3 = require('web3');
        let web3;
        //当前存在web3实例
        if (typeof window.web3 !== 'undefined') {
            web3 = new Web3(window.web3.currentProvider)
        }
        let from = this.props.token;
        var ethUtil = require('ethereumjs-util');
        var msg = ethUtil.bufferToHex(new Buffer('forever bird'),'utf8');
        var params = [msg, from];
        var method = 'personal_sign';
        web3.currentProvider.sendAsync({
          method,
          params,
          from},(err,result) => {
            console.log(err);
            console.log(result);
            if(result.result){
              let data = {
                method : 'POST',
                credentials: 'same-origin',
                headers: {
                  'Content-Type': 'text/plain'
                },
                body : JSON.stringify({
                  email : values.email,
                  nick : values.nickname,
                  address : this.props.token,
                  sign : result.result
                }),
            }
            fetch(BASE_URL+'/api/register',data)
            .then(response => {
            return response.json()
            }).then(json => {
              if(json.status === 0){
                this.setState({
                  redirect : true
              });
              }else{
                alert('regeist error.error msg is ' + json.message);
              }
            }) 
            }
          }
        );
        /* web3.eth.sign(from,web3.sha3(from+values.email+values.nickname),(err,res) => {
            if(err == null){
              let data = {
                  method : 'POST',
                  credentials: 'same-origin',
                  // mode : 'cors',
                  headers: {
                    'Content-Type': 'text/plain'
                  },
                  body : JSON.stringify({
                    email : values.email,
                    nick : values.nickname,
                    address : this.props.token
                  }),
              }
              fetch(BASE_URL+'/register',data)
              .then(response => {
              return response.json()
              }).then(json => {
                if(json.status === 0){
                  this.setState({
                    redirect : true
                });
                }else{
                  alert('regeist error.error msg is ' + json.message);
                }
              }) 
            }
        }); */
      }
    });
  }

  handleConfirmBlur = (e) => {
    const value = e.target.value;
    this.setState({ confirmDirty: this.state.confirmDirty || !!value });
  }

  render() {
    const FormItem = Form.Item;
    const { getFieldDecorator } = this.props.form;

    const formItemLayout = {
      labelCol: {
        xs: { span: 24 },
        sm: { span: 8 },
      },
      wrapperCol: {
        xs: { span: 24 },
        sm: { span: 16 },
      },
    };
    const tailFormItemLayout = {
      wrapperCol: {
        xs: {
          span: 24,
          offset: 0,
        },
        sm: {
          span: 24,
          offset: 10,
        },
      },
    };

    if(this.state.redirect){
        return(
            <Redirect to = {{
                    pathname : '/mybird',
                    state : {
                        token :this.props.token,
                    }
                }}
            />
        )
    }else{
        const email = this.props.email;
        const nick = this.props.nick;
        return (
          <Form onSubmit={this.handleSubmit}>
            <FormItem
              {...formItemLayout}
              label="钱包地址"
            >
                <Input defaultValue={this.props.token} disabled/>
            </FormItem>
            {email ?
                <FormItem
                {...formItemLayout}
                label="E-mail"
                >
                <Input defaultValue={email} disabled/>
                </FormItem>
               :
              <FormItem
                {...formItemLayout}
                label="E-mail"
              >
                {getFieldDecorator('email', {
                  rules: [{
                    type: 'email', message: 'The input is not valid E-mail!',
                  }, {
                    required: true, message: 'Please input your E-mail!',
                  }],
                })(
                  <Input />
                )}
              </FormItem>
            }
            {nick ?
                <FormItem
                {...formItemLayout}
                label="Nickname"
                >
                <Input defaultValue={nick} disabled/>
                </FormItem>
               :
              <FormItem
                {...formItemLayout}
                label={(
                  <span>
                    Nickname&nbsp;
                    <Tooltip title="What do you want others to call you?">
                      <Icon type="question-circle-o" />
                    </Tooltip>
                  </span>
                )}
              >
                {getFieldDecorator('nickname', {
                  rules: [{ required: true, message: 'Please input your nickname!', whitespace: true }],
                })(
                  <Input />
                )}
              </FormItem>
            }
            {email && nick ?
                <div></div>
               :
              <FormItem {...tailFormItemLayout}>
                <Button type="primary" htmlType="submit">Register</Button>
              </FormItem>
            }
          </Form>
        );
    }
  }
}
