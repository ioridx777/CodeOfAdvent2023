package main

import (
	"fmt"
	"os"
	// "regexp"
	// "strconv"
	"strings"
)

var expandedCol=[]int{}
var expandedRow=[]int{}

var markedPos=[][2]int{}

func min(a int,b int)int{
  if a<b{
  return a
  } else{
    return b
  }
}
func max(a int,b int)int{
    if a>b{
  return a
  } else{
    return b
  }
}

func addRowLen(yStart int,yEnd int)int{
  result :=0
  for _,eCol:=range expandedRow{
    if eCol>yStart&&eCol<yEnd{
      result++
    }
  }
  return result
}
func addColLen(xStart int,xEnd int)int{
  result :=0
  for _,eCol:=range expandedCol{
    if eCol>xStart&&eCol<xEnd{
      result++
    }
  }
  return result
}

func getDist(posA [2]int,posB[2]int)int{
  xStart:=min(posA[0],posB[0])
  yStart:=min(posA[1],posB[1])
  xEnd:=max(posA[0],posB[0])
  yEnd:=max(posA[1],posB[1])
  addedLength:=addColLen(xStart,xEnd)
  addedLength+=addRowLen(yStart,yEnd)

  // fmt.Println(`Here with `,posA,` and `,posB,">",addedLength)
  return xEnd-xStart+yEnd-yStart+addedLength*999999
}

func getSumPos(currentPos [2]int,otherPos [][2]int)int{
  sum:=0
  for _,pos:=range otherPos{
    sum+=getDist(currentPos,pos)
  }
  return sum
}

func main() {
  inputBytes, error := os.ReadFile("inputDay11.txt")

	if error != nil {
		fmt.Println("fail")
		return
	}
	inputText := string(inputBytes)

	linesOfText := strings.Split(inputText, "\n")
  width := len(linesOfText[0])
	for rowId, line := range linesOfText {
		fmt.Println(line)
		if len(line) == 0 {
			continue
		}
     isExpandRow:=true
    for i:=0;i<len(line);i++{
      if(line[i:i+1]!="."){
        isExpandRow=false
        break
      }
    }
    if isExpandRow{
      expandedRow = append(expandedRow,rowId )
    }
	}

  for i:=0;i<width;i++{
    isExpandCol:=true
    for j:=0;j<len(linesOfText)-1;j++{
      fmt.Println(linesOfText[j])
      if linesOfText[j][i:i+1]!="."{
        isExpandCol=false
        markedPos=append(markedPos, [2]int{i,j})
      } 
    }
    if isExpandCol{
      expandedCol=append(expandedCol, i)
    }
  }
  fmt.Println(`expanded row `,expandedRow)
  fmt.Println(`expanded col `,expandedCol)
  fmt.Println(`markedpos `,markedPos)
  sum:=0
  for posId,marked:=range markedPos{
    if posId+1>=len(markedPos){
      break
    }
    remainingPos:=markedPos[posId+1:]
    
    sum+=getSumPos(marked,remainingPos)
  }
  fmt.Println(sum)
	fmt.Println(`hellowlrd`)
}
