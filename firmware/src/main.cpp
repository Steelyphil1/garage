#include <WiFi.h>
#include <WiFiClientSecure.h>
#include <HTTPClient.h>

const char* ssid = WIFI_NAME;
const char* password = WIFI_PASSWORD;

const int reedPin = 4;
const int ledPin = 2;

bool lastState = HIGH;

void sendEvent(String state) {
  if (WiFi.status() != WL_CONNECTED) {
    Serial.println("WiFi not connected");
    return;
  }

  WiFiClientSecure client;
  client.setInsecure();

  HTTPClient http;

  Serial.println("Starting HTTPS request...");

  http.begin(client, GARAGE_ENDPOINT);
  http.addHeader("Content-Type", "application/json");

  String payload = "{\"state\":\"" + state + "\"}";

  Serial.print("Sending: ");
  Serial.println(payload);

  int code = http.PUT(payload);

  Serial.print("HTTP code: ");
  Serial.println(code);

  if (code > 0) {
    Serial.println(http.getString());
  }

  http.end();
}

void setup() {
  pinMode(reedPin, INPUT_PULLUP);
  pinMode(ledPin, OUTPUT);

  Serial.begin(115200);
  delay(2000);

  Serial.println("Booting ESP32...");

  WiFi.begin(ssid, password);

  while (WiFi.status() != WL_CONNECTED) {
    delay(500);
    Serial.print(".");
  }

  Serial.println("WiFi connected!");
  Serial.println(WiFi.localIP());

  lastState = digitalRead(reedPin);
}

void loop() {
  int state = digitalRead(reedPin);

  digitalWrite(ledPin, state == LOW ? HIGH : LOW);

  if (state != lastState) {
    delay(50);

    if (state == LOW) {
      Serial.println("Door CLOSED");
      sendEvent("Closed");
    } else {
      Serial.println("Door OPEN");
      sendEvent("Open");
    }

    lastState = state;
  }

  delay(50);
}