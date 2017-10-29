//Zhenyu Yang 2017/10/27
// microbals competition: Chthamalus/Balanus,
func DrawPetri(width int) Canvas{
   white:= MakeColor(255,255,255)
   black := MakeColor(0, 0, 0)
   pic:= CreateNewCanvas(width,width)
   pic.SetFillColor(black)
   pic.Clear()
   pic.SetFillColor(white)
   pic.Circle(width/2,width/2,width/2)
   return pic
}

func DrawBacteria(x,y,r float64, pic Canvas, color color.Color) Canvas{
   pic.SetFillColor(color)
   pic.Circle(x,y,r)
}

func InitialPosition(width int,n int,pic Canvas) Canvas {
    blue := MakeColor(0,0,225)
    rand.Seed(time.Now().UTC().UnixNano())
    for i:=0 ; i < n ;i++{
      x := rand.Intn(width)
      y := rand.Intn(width)
      DrawBacteria(x,y,r,pic,blue)
      }
}
// Control with temperature
func DrawHotDeadCell(x,y,r float64, pic Canvas) Canvas{ // if the temperature is too hot
  // cell will dissolve its self
    m := 0
    n := 0
    for i:=0; i < 5; i++{
    color := MakeColor(m, n, 255)
    DrawBacteria(x,y,r,pic,color)
    m = m + 51
    n = n + 51
  } // hot to death cell will turn into white
  return pic
}

func DrawFreezeDeadCell(x,y,r float64, pic Canvas) Canvas{ // if the temperature is too hot
  // cell will dissolve its self
    m := 0
    n := 255
    for i := 1; i <= 2 ; i++{
    color := MakeColor(m, 0, n)
    DrawBacteria(x,y,r,pic,color)
    m = m + 64
    n = n - 64
  }
  return pic
}

func DrawAboveBestTemp(x,y,r int, speed,currenttemp, besttemp,hotdeadtemp float64, pic Canvas) Canvas{
  // when the environment's temperature we above best temp

   if currenttemp > besttemp && currenttemp < hotdeadtemp {
     speed = speed + speed * (currenttemp - besttemp)/10
   }
   if currenttemp > hotdeadtemp {
      speed = 0
      DrawHotDeadCell(x,y,r,pic)
   }
   return pic
}

func DrawBelowBestTemp(x,y,r int, maxspeed,speed,currenttemp,besttemp, freezedeadtemp float64, pic Canvas) Canvas{
    // when the environment's temperature we below best temp
    blue := MakeColor(0,0,255)
    if currenttemp < besttemp && currenttemp > 0 {
      speed = maxspeed * currenttemp / besttemp
      DrawBacteria(x,y,r,blue)
    }
    if currenttemp < = 0 {
      speed = 0
      DrawColdDeadCell(x,y,r,pic)
    }
    return pic
}

// Control with oxygen
func IsAnaerobic(bactname string) bool {
  if bactname == ""{
    return true
  }
  return false
}

func DrawAboveBestOxy(x,y,r int, oxy,bestoxy,speed,destroyoxy float64,name string){
  if IsAnaerobic(name) == false && oxy < destroyoxy {
    speed = speed * oxygen/bestoxy
  } else if IsAnaerobic(name) == true && oxy > 0{
    speed = 0
  } else if
}


func ChangeVitamin(){}

func ChangeMineral(){}

func ChangeAgarVisible(){}
func ChangeOxygen(){}
func ChangePetones(){}
func ChangeLight(){}
