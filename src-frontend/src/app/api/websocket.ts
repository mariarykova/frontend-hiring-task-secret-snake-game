export const connectToWebsocket = (setBid: any) => {
  const ws = new WebSocket("ws://localhost:8080/api/updates");

  ws.onopen = () => {
    console.log("WebSocket connection established.");
  };

  ws.onmessage = (event) => {
    const message = event.data;
    console.log("Received message:", JSON.parse(message));
    setBid(JSON.parse(message));
  };

  ws.onclose = () => {
    console.log("WebSocket connection closed.");
  };

  ws.onerror = (error) => {
    console.error("WebSocket error:", error);
  };

  return () => {
    ws.close();
  };
};
