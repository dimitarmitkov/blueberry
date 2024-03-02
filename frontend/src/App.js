// src/App.js
import React, { useState, useEffect } from 'react';
import { Button, Row, List, Input } from 'antd';
import { SendOutlined } from '@ant-design/icons';
import './App.css';


function App() {
  const [messages, setMessages] = useState([]);
  const [newMessage, setNewMessage] = useState('');

  useEffect(() => {

    fetch('/api')
        .then(response => response.json())
        .then(data => setMessages(data))
        .catch(error => console.error('Error fetching data:', error));
  }, []);

  const handleInputChange = (event) => {
    setNewMessage(event.target.value);
  };

  const handleSubmit = (event) => {
    event.preventDefault();

    // Send a POST request to add a new message
    fetch('/api', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ content: newMessage }),
    })
        .then(response => response.json())
        .then(data => {
          // Refresh the list of messages after adding a new one
          setMessages([...messages, data]);
          setNewMessage('');
        })
        .catch(error => console.error('Error adding new message:', error));
  };

  const handleClick = ()=>{
    window.location.reload()
  }

  return (
      <>
        <div className="App">
          <header className="App-header">
            <h1>Messages from the backend:</h1>
            <List
                dataSource={messages}
                renderItem={message => (
                    <List.Item><a href={message.Content} target="_blank">{message.Link}</a></List.Item>
                )}
            />

            <form onSubmit={handleSubmit} style={{marginTop: '20px'}}>
              <Input
                  type="text"
                  value={newMessage}
                  onChange={handleInputChange}
                  placeholder="New Message"
              />
              <Button
                  type="primary"
                  icon={<SendOutlined/>}
                  style={{marginLeft: '10px'}}
                  onClick={handleSubmit}
              >
                Add Message
              </Button>
            </form>
          </header>
          <Row>
            <Button type="primary" size="large" onClick={handleClick}>Refresh</Button>
          </Row>
        </div>
      </>
  );
}


export default App;