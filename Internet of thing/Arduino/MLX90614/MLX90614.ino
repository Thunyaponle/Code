47
#include<TFT_eSPI.h>
#include<Adafruit_MLX90614.h>
#include<Wire.h>
TFT_eSPI tft;

Adafruit_MLX90614 mlx = Adafruit_MLX90614();
void setup() {
  mlx.begin();
  tft.begin();
  tft.setRotation(3);
  tft.setTextSize(3);
  tft.fillScreen(TFT_BLACK);
  tft.setTextColor(TFT_WHITE,TFT_BLACK);
  tft.drawString("MLX90614 ",50,30);
  tft.drawString("IR-TEMPERATURE",50,50);
  tft.setTextSize(2);

}

void loop() {
  float ambTempC = mlx.readAmbientTempC();
  float objTempC = mlx.readObjectTempC();
  tft.drawString(String("Ambient: ")+String(ambTempC)+String(" C "),50,120);
  tft.drawString(String("Object: ")+String(objTempC)+String(" C "),50,150);
}
