import React, {useState} from 'react';
import {Input, Button, Select} from 'antd';
import { Paper, Typography } from "@material-ui/core";
import {Option} from "antd/lib/mentions";

const Chaincode = () =>{
    const [action,setAction] = useState('');
    const [channelId,setChannelId] = useState('mychannel');
    const [result,setResult] = useState('');
    const [name,setName] = useState('simplecc');
    const [version,setVersion] = useState('1.0.0');
    const [sequence,setSequence] = useState(1);
    const [initRequired,setInitRequired] = useState(true);

    const onPackage = () => {
        const myHeaders = new Headers();
        myHeaders.append("Content-Type", "application/json");    
        const raw = JSON.stringify({
            "chaincode_id": name,
            "chaincode_version" : version
        });
        const requestOptions = {
            method: 'POST',
            headers: myHeaders,
            body: raw,
            redirect: 'follow'
        };
        fetch("http://localhost:8080/fabric/chaincodes/package", requestOptions)
            .then(response => {
                return response.json()
            })
            .then((json) => {
                console.log(json)
                setResult(json.message)
            })
            .catch(error => {
                console.log('error', error)
                setResult(result)
            });
    }
    const onInstall = () => {
        console.log(version)
        const myHeaders = new Headers();
        myHeaders.append("Content-Type", "application/json");    
        const raw = JSON.stringify({
            "chaincode_id": name,
            "chaincode_version" : version
        });
        const requestOptions = {
            method: 'POST',
            headers: myHeaders,
            body: raw,
            redirect: 'follow'
        };
        fetch("http://localhost:8080/fabric/chaincodes/install", requestOptions)
            .then(response => {
                return response.json()
            })
            .then((json) => {
                console.log(json)
                setResult(json.message)
            })
            .catch(error => {
                console.log('error', error)
                setResult(result)
            });
    }
    const onGetInstall = () => {
        fetch("http://localhost:8080/fabric/chaincodes/install")
            .then(response => {
                return response.json()
            })
            .then((json) => {
                console.log(json)
                setResult(json.message)
            })
            .catch(error => {
                console.log('error', error)
                setResult(result)
            });
    }
    const onApprove = () => {
        const myHeaders = new Headers();
        myHeaders.append("Content-Type", "application/json");
        const raw = JSON.stringify({
            "chaincode_id": name,
            "chaincode_version" : version,
            "sequence" : sequence,
            "init_required" : initRequired
        });
        const requestOptions = {
            method: 'POST',
            headers: myHeaders,
            body: raw,
            redirect: 'follow'
        };
        fetch("http://localhost:8080/fabric/chaincodes/approve", requestOptions)
            .then(response => {
                return response.json()
            })
            .then((json) => {
                console.log(json)
                setResult(json.message)
            })
            .catch(error => {
                console.log('error', error)
                setResult(result)
            });
    }
    const onGetApprove = () => {
        fetch("http://localhost:8080/fabric/chaincodes/approve")
            .then(response => {
                return response.json()
            })
            .then((json) => {
                console.log(json)
                setResult(json.message)
            })
            .catch(error => {
                console.log('error', error)
                setResult(result)
            });
    }
    const onCommit = () => {
        const myHeaders = new Headers();
        myHeaders.append("Content-Type", "application/json");
        const raw = JSON.stringify({
            "chaincode_id": name,
            "chaincode_version" : version,
            "sequence" : sequence,
            "init_required" : initRequired
        });
        const requestOptions = {
            method: 'POST',
            headers: myHeaders,
            body: raw,
            redirect: 'follow'
        };
        fetch("http://localhost:8080/fabric/chaincodes/commit", requestOptions)
            .then(response => {
                return response.json()
            })
            .then((json) => {
                console.log(json)
                setResult(json.message)
            })
            .catch(error => {
                console.log('error', error)
                setResult(result)
            });
    }
    const onGetCommit = () => {
        fetch("http://localhost:8080/fabric/chaincodes/commit")
            .then(response => {
                return response.json()
            })
            .then((json) => {
                console.log(json)
                setResult(json.message)
            })
            .catch(error => {
                console.log('error', error)
                setResult(result)
            });
    }
    const onInit = () => {
        const myHeaders = new Headers();
        myHeaders.append("Content-Type", "application/json");
        const raw = JSON.stringify({
            "chaincode_id": name,
            "channel_name" : channelId,
            "init_required" : initRequired
        });
        const requestOptions = {
            method: 'POST',
            headers: myHeaders,
            body: raw,
            redirect: 'follow'
        };
        fetch("http://localhost:8080/fabric/chaincodes/init", requestOptions)
            .then(response => {
            return response.json()
        })
            .then((json) => {
                console.log(json)
                setResult(json.message)
            })
            .catch(error => {
                console.log('error', error)
                setResult(result)
            });
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
        console.log(initRequired)
        setInitRequired(e.target.value);
    };
    const onChangeChannelId = (e) => {
        console.log(channelId)
        setChannelId(e.target.value);
    };
    const onChangeAction = (value) => {
        setAction(value)
    }
    const onSubmit =(e) => {
        switch (action) {
            case 'package' :
                onPackage(e);
                break;
            case 'install' :
                onInstall(e);
                break;
            case 'getInstall' :
                onGetInstall(e);
                break;
            case 'approve' :
                onApprove(e);
                break;
            case 'getApprove' :
                onGetApprove(e);
                break;
            case 'commit' :
                onCommit(e);
                break;
            case 'getCommit' :
                onGetCommit(e);
                break;
            case 'init' :
                onInit(e);
                break;
            default:
                setResult("invalid action selcted")
        }
    }

    return (
        <>
            <br/><br/><br/><br/><br/><br/><br/><br/><br/>
            <Paper style={{ width: "50%" ,  margin: "0 auto"}}>
        <Typography variant='h5'>체인코드 관리</Typography>

        <form  style={{padding:10}}>
            <div>
                <label htmlFor="channel_id">채널 이름</label><br/>
                <Input name="channel_id" value={channelId} required onChange={onChangeChannelId} />
            </div>
            <div>
                <label htmlFor="chaincode_name">체인코드 이름</label><br/>
                <Input name="chaincode_name" value={name} required onChange={onChangeName} />
            </div>
            <div>
                <label htmlFor="version">체인코드 버전</label><br/>
                <Input name="version" value={version} required onChange={onChangeVersion} />
            </div>
            <div>
                <label htmlFor="sequence">sequence</label><br/>
                <Input name="sequence"  value={sequence} required onChange={onChangeSequence} />
            </div>
            <div>
                <label htmlFor="init_required">InitRequired</label><br/>
                <Input name="init_required" value={initRequired} required onChange={onChangeInitRequired} />
            </div>
            <div>
                <label htmlFor="action">동작</label><br/>
                <Select onChange={onChangeAction} style={{width : 200}}>
                    <Option value="package">패키징</Option>
                    <Option value="install">설치</Option>
                    <Option value="getInstall">설치된 정보</Option>
                    <Option value="approve">승인</Option>
                    <Option value="getApprove">승인된 정보</Option>
                    <Option value="commit">커밋</Option>
                    <Option value="getCommit">커밋된 정보</Option>
                    <Option value="init">생성</Option>
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
            <Typography variant='h6'>응답 </Typography>
            <div>
                {result}
            </div>
        </Paper>
        </>
    );
};

export default Chaincode;