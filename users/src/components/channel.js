import React, {useState} from 'react';
import { Input , Button} from 'antd';
import { Paper, Typography } from "@material-ui/core";

const Channel = () =>{
    const [name,setName] = useState('');
    const [result,setResult] = useState('');
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
                console.log(response.body)
                setResult(response.body.result)
            }
                )
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
                console.log(response.body)
                setResult(response.body.result)
            }
                )
            .catch(error => console.log('error', error));
    };


    // Coustom Hook 이전
    const onChangeName = (e) => {
        setName(e.target.value);
    };


    return (
        <>
            <Paper style={{ width: "50%" ,  margin: "0 auto"}}>
            <Typography variant='h5'>채널 설정</Typography>

            <form  style={{padding:10}}>
                <div>
                    <label htmlFor="channel_name">채널 이름</label><br/>
                    <Input name="channel_name" value={name} required onChange={onChangeName} />
                </div>
                <div style={{marginTop:10}}>
                    <Button type="primary" onClick={createClick} >생성</Button>
                    <br></br>
                    <br></br>
                    <Button type="primary" onClick={joinClick} >참여</Button>
                </div>
            </form>
            </Paper>
            <br></br>
            <br></br>
            <br></br>
            <Paper style={{ width: "50%" ,  margin: "0 auto"}}>
            <Typography variant='h6'>응답 </Typography>
            <div>
                {result}
            </div>
            </Paper>
        </>
        
    );
};

export default Channel;