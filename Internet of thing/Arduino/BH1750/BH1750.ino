#include<TFT_eSPI.h>
#include"BH1750.h"
#include<Wire.h>
BH1750 lightMeter;

TFT_eSPI tft;

void setup() {
  Wire.begin();
  lightMeter.begin();
  tft.begin();
  tft.setRotation(3);
  tft.setTextSize(3);
  tft.fillScreen(TFT_BLACK);
  tft.setTextColor(TFT_WHITE,TFT_BLACK);
  tft.drawString("BH1750 ",50,30);
  tft.drawString("Lux meter",50,50);
  tft.setTextSize(2);
}

void loop() {
  int lux = (int)lightMeter.readLightLevel();
  tft.drawString(String("Light: ")+String(lux)+String("lx "),50,120);
}
