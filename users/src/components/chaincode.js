import React, {useState} from 'react';
import { Input  , Button} from 'antd';
import { Paper, Typography } from "@material-ui/core";

const Chaincode = () =>{
    const [result,setResult] = useState('');
    const [name,setName] = useState('');
    const [version,setVersion] = useState('');
    const [sequence,setSequence] = useState('');
    const [initRequired,setInitRequired] = useState('');

    const onPackage = () => {
        const myHeaders = new Headers();
        myHeaders.append("Content-Type", "application/json");    
        const raw = JSON.stringify({
            "email": id,
            "name" : nick,
            "pwd": password
        });
        const requestOptions = {
            method: 'POST',
            headers: myHeaders,
            body: raw,
            redirect: 'follow'
        };
        fetch("http://localhost:8080/fabric/chaincodes/install", requestOptions)
            .then(response => response.text())
            .then(result => {
                console.log("result : "+result)
                setResult(result)
            })
            .catch(error => {
                console.log('error', error)
                setResult(result)
            });
    }
    const onInstall = () => {
        const myHeaders = new Headers();
        myHeaders.append("Content-Type", "application/json");    
        const raw = JSON.stringify({
            "email": id,
            "name" : nick,
            "pwd": password
        });
        const requestOptions = {
            method: 'POST',
            headers: myHeaders,
            body: raw,
            redirect: 'follow'
        };
        fetch("http://localhost:8080/fabric/chaincodes/install", requestOptions)
            .then(response => response.text())
            .then(result => {
                console.log("result : "+result)
                setResult(result)
            })
            .catch(error => {
                console.log('error', error)
                setResult(result)
            });
    }
    const onGetInstall = () => {

    }
    const onApprove = () => {

    }
    const onGetApprove = () => {

    }
    const onCommit = () => {

    }
    const onGetCommit = () => {

    }
    const onInit = () => {

    }
    const onChangeName = (e) => {
        setName(e.target.value);
    };
    const onChangeVersion = (e) => {
        setVersion(e.target.value);
    };
    const onChangeSequence = (e) => {
        setSequence(e.target.value);
    };
    const onChangeInitRequired = (e) => {
        setInitRequired(e.target.value);
    };

    return (
        <>
            <Paper style={{ width: "50%" ,  margin: "0 auto"}}>
        <Typography variant='h5'>체인코드 관리</Typography>

        <form style={{padding:10}}>
            <div>
                <label htmlFor="chaincode_name">체인코드 이름</label><br/>
                <Input name="chaincode_name" value={name} required onChange={onChangeName} />
            </div>
            <div>
                <label htmlFor="name">체인코드 버전</label><br/>
                <Input name="name" value={version} required onChange={onChangeVersion} />
            </div>
            <div>
                <label htmlFor="sequence">sequence</label><br/>
                <Input name="sequence"  value={sequence} required onChange={onChangeSequence} />
            </div>
            <div>
                <label htmlFor="init_required">InitRequired</label><br/>
                <Input name="init_required" value={initRequired} required onChange={onChangeInitRequired} />
            </div>
            <div style={{marginTop:10}}>
                <Button type="primary" style={{margin: 10}} onClick={onPackage} >package</Button>
                <br></br>
                <Button type="primary" style={{margin: 10}} onClick={onInstall} >install</Button>
                <Button type="primary" style={{margin: 10}} onClick={onGetInstall} >get install</Button>
                <br></br>
                <Button type="primary" style={{margin: 10}} onClick={onApprove} >approve</Button>
                <Button type="primary" style={{margin: 10}} onClick={onGetApprove} >get approve</Button>
                <br></br>
                <Button type="primary" style={{margin: 10}} onClick={onCommit} >commit</Button>
                <Button type="primary" style={{margin: 10}} onClick={onGetCommit} >get commit</Button>
                <br></br>
                <Button type="primary" style={{margin: 10}} onClick={onInit} >init</Button>
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

export default Chaincode;