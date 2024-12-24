#include <TFT_eSPI.h>
TFT_eSPI tft;
#define LCD_BACKLIGHT (72Ul)
void setup() {
  // put your setup code here, to run once:
  tft.begin();
  tft.fillScreen(TFT_RED);
}

void loop() {
  // put your main code here, to run repeatedly:
  digitalWrite(LCD_BACKLIGHT, HIGH);
  delay(1000);
  digitalWrite(LCD_BACKLIGHT, LOW);
  delay(1000);