#include <TFT_eSPI.h>
TFT_eSPI tft;

void setup() {
  // put your setup code here, to run once:
  tft.begin();
  tft.drawString("Hello world!",0,0);
}

void loop() {
  // put your main code here, to run repeatedly:

}
