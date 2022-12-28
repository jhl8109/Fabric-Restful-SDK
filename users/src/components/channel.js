import React, {useState} from 'react';
import { Select, Input , Button} from 'antd';
import { Paper, Typography } from "@material-ui/core";
import {Option} from "antd/lib/mentions";

const Channel = () =>{
    const [name,setName] = useState('');
    const [result,setResult] = useState('');
    const [action,setAction] = useState('');
    const createClick = (e) => {
        e.preventDefault();
        const myHeaders = new Headers();
        myHeaders.append("Content-Type", "application/json");

    
        const raw = JSON.stringify({
            "channel_name": name
        });
    
        const requestOptions = {
            method: 'POST',
            headers: myHeaders,
            body: raw,
            redirect: 'follow'
        };
    
        fetch("http://localhost:8080/fabric/channels/create", requestOptions)
            .then(response => {
                return response.json()
            })
            .then((json) => {
                console.log(json)
                setResult(json.message)
            })
            .catch(error => console.log('error', error));
    };
    const joinClick = (e) => {
        e.preventDefault();
        const myHeaders = new Headers();
        myHeaders.append("Content-Type", "application/json");

    
        const raw = JSON.stringify({
            "channel_name": name
        });
    
        const requestOptions = {
            method: 'POST',
            headers: myHeaders,
            body: raw,
            redirect: 'follow'
        };
    
        fetch("http://localhost:8080/fabric/channels/join", requestOptions)
            .then(response => {
                return response.json()
            })
            .then((json) => {
                console.log(json)
                setResult(json.message)
            })
            .catch(error => console.log('error', error));
    };


    // Coustom Hook 이전
    const onChangeName = (e) => {
        setName(e.target.value);
    };
    const onChangeAction = (value) => {
        setAction(value)
    }
    const onSubmit = (e) => {
        switch(action) {
            case 'create' :
                createClick(e)
                break
            case 'join' :
                joinClick(e)
                break
            default :
                setResult("invalid action selected")
        }
    }

    return (
        <>
            <br/><br/><br/><br/><br/><br/><br/><br/><br/>
            <Paper style={{ width: "50%" ,  margin: "0 auto"}}>
            <Typography variant='h5'>채널 설정</Typography>

            <form  style={{padding:10}}>
                <div>
                    <label htmlFor="channel_name">채널 이름</label><br/>
                    <Input name="channel_name" value={name} required onChange={onChangeName} />
                </div>
                <div>
                    <label htmlFor="action">동작</label><br/>
                    <Select onChange={onChangeAction} style={{width : 100}}>
                        <Option value="create">생성</Option>
                        <Option value="join">참여</Option>
                    </Select>
                </div>
                <div style={{ margin: "0 auto" ,marginTop:10 }}>
                    <Button type="primary" onClick={onSubmit}>제출</Button>
                </div>
            </form>
            </Paper>
            <br></br>
            <br></br>
            <br></br>
            <Paper style={{ width: "50%" ,  margin: "0 auto"}}>
            <Typography variant='h6'>응답</Typography>
            <div>
                {result}
            </div>
            </Paper>
        </>
        
    );
};

export default Channel;