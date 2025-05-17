import React, { useState } from 'react';
import './App.css';

function App() {
  const [agentID, setAgentID] = useState('');
  const [command, setCommand] = useState('');
  const [response, setResponse] = useState('');
  const [loading, setLoading] = useState(false);

  const handleSend = async () => {
    if (!agentID || !command) {
      alert("Please enter both agent ID and command.");
      return;
    }

    setLoading(true);
    setResponse('');

    try {
      // Send command
      await fetch('http://localhost:8080/send', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ agent_id: agentID, command })
      });

      // Poll response every second up to 10 seconds
      for (let i = 0; i < 10; i++) {
        const res = await fetch(`http://localhost:8080/response/${agentID}`);
        if (res.ok) {
          const data = await res.json();
          setResponse(data.response);
          break;
        }
        await new Promise(resolve => setTimeout(resolve, 1000));
      }

    } catch (err) {
      console.error("Error:", err);
      setResponse("❌ Error occurred while sending command or fetching response.");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="App">
      <h2>💻 C2 Command Interface</h2>
      <input
        type="text"
        placeholder="Agent ID"
        value={agentID}
        onChange={e => setAgentID(e.target.value)}
      /><br />
      <input
        type="text"
        placeholder="Command (e.g., whoami)"
        value={command}
        onChange={e => setCommand(e.target.value)}
      /><br />
      <button onClick={handleSend} disabled={loading}>
        {loading ? "Sending..." : "Send Command"}
      </button>

      {response && (
        <div className="response">
          <h4>📬 Response:</h4>
          <pre>{response}</pre>
        </div>
      )}
    </div>
  );
}

export default App;
