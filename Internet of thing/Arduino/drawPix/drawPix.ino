#include <TFT_eSPI.h>
TFT_eSPI tft;

void setup() {
  // put your setup code here, to run once:
  tft.begin();
  tft.setRotation(3);
  tft.fillScreen(TFT_BLUE);
  tft.drawPixel(4,7,TFT_YELLOW);
}

void loop() {
  // put your main code here, to run repeatedly:

}