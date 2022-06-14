function overlap(v)
{
     
    // Variable to store the maximum
    // count
    var ans = 0;
    var count = 0;
    var data = [];
 
    // Storing the x and y
    // coordinates in data vector
    for(var i = 0; i < v.length; i++)
    {
         
        // Pushing the x coordinate
        data.push([v[i][0], 'x']);
 
        // Pushing the y coordinate
        data.push([v[i][1], 'y']);
    }
 
    // Sorting of ranges
  data.sort((a,b) => a[0] - b[0]);
  console.log('sorted ', data)
 
    // Traverse the data vector to
    // count number of overlaps
    for(var i = 0; i < data.length; i++)
    {
         
        // If x occur it means a new range
        // is added so we increase count
        if (data[i][1] == 'x')
            count++;
 
        // If y occur it means a range
        // is ended so we decrease count
        if (data[i][1] == 'y')
            count--;
 
        // Updating the value of ans
        // after every traversal
        ans = Math.max(ans, count);
    }
  return ans;
 
    // Printing the maximum value
}

function solve_1() {
  var myjson = `
[
    ["User_001", 1, 3],
    ["User_001", 4, 6],
    ["User_001", 2, 5]
]  `;
  var data = JSON.parse(myjson);
  if (data.length <=0) {
    return 0;    
  }

  var curMap = [];
  curMap.push([...data[0], 1])
  data.sort((a,b) => a[1] - b[1])
  ans = 0;

  const prepareData =  data.map(a => [a[1],a[2]]);
  console.log('the maximum ' + overlap(prepareData))

  // var listOfTime = [];
  // for(let i = 0;i<data.length;i++) {
  //   listOfTime.push({time: data[i][1], type: 'start'});
  //   listOfTime.push({time: data[i][2], type: 'end'});
  // }
  //
  // listOfTime.sort((a,b) => a.time - b.time);
  // console.log(listOfTime);
  //
  // let count = 0;
  // for(let i = 0; i < listOfTime.length;i++) {
  //   if(listOfTime[i].type === 'start') count++;
  //   else count--;
  //
  //   ans = Math.max(ans, count);
  // }
  // console.log(`The mamximum number of concurrent stream is ${ans}`);

  console.log(data)
  for(let i = 1;i<data.length;i++) {
    const cur = data[i];
    let isOverlap = false;
    for(let j = 0; j <curMap.length;j++) {
      const curInterval = curMap[j];
      if(
        cur[2] >= curInterval[1] && cur[2] < curInterval[2]
        || cur[1] >= curInterval[1] && cur[1] <= curInterval[2]
        || cur[1] <= curInterval[1] && cur[2] >= curInterval[2]) {
        curInterval[1] = Math.min(curInterval[1], cur[1]);
        curInterval[2] = Math.max(curInterval[2], cur[2]);
        curInterval[3]++;
        ans = Math.max(ans, curInterval[3]);
        isOverlap = true;
        break;
      }    
    }
    if(!isOverlap) {
      curMap.push([...cur, 1])
    }
  }
  console.log(curMap)

  console.log(`The mamximum number of concurrent stream is ${ans}`);

}
// solve_1()

function solve_2() {
   var myjson = `
{
  "Learner-0001": [
    "Course-0001",
    "Course-0002",
    "Course-0003"
  ],
  "Learner-0002": [
    "Course-0002",
    "Course-0003",
    "Course-0004"
  ],
  "Learner-0003": [
    "Course-0004",
    "Course-0005",
    "Course-0006"
  ],
  "Learner-0004": [
    "Course-0005",
    "Course-0006",
    "Course-0007"
  ]
}`; 
  const data = JSON.parse(myjson);
  const map = {};
  Object.keys(data).forEach(key => {
    const set = new Set(data[key]);
    set.forEach(a => { 
      if(map[a]) {
        map[a]++;
      } else {
        map[a] = 1;
      }
    })
    })

  const ans = Object.keys(map).filter(key => {
    return map[key] === 1
  });

  console.log(ans);
}

// solve_2()

function solve_3() {
   var answers = ["A", "B", "A", "C", "D"]; 
  var learners = [
    ["A", "B", "B", "C", "D"],
    ["C", "B", "A", "D", "B"],
    ["A", "B", "C", "D", "C"],
    ["B", "B", "A", "D", "A"],
    ["A", "B", "D", "D", "D"]
  ]
  var counters = {};
  answers.forEach((a,index) => {
    counters[index] = 0;
  });
  var ans = 0;
  var ansIndex = 0;

  for(let i =0;i < learners.length;i++) {
    for (let j = 0; j < answers.length;j++) {
      if(learners[i][j] === answers[j]) {
        counters[j]++;
        if(ans < counters[j]) {
          ans =counters[j]; 
          ansIndex = j;
        }
      }
    }
  }
  console.log(counters)
  console.log("The easiest question is index " + ansIndex)
  return ans;

}

// solve_3()
function solve_4() {
  const courseOrders =[
    ["Course_001","Course_002","Course_001","Course_003"],
    ["Course_002","Course_001","Course_001","Course_004"]
] 
  const counters = {}

  for(let i = 0;i < courseOrders.length;i++) {
    for(let j = 0;j<courseOrders[i].length-1;j++) {
      const currentCourse = courseOrders[i][j];
      const currentNextCourse = courseOrders[i][j+1];
      if(counters[currentCourse]) {
        if(counters[currentCourse][currentNextCourse]) {
          counters[currentCourse][currentNextCourse]++;
        } else {
          counters[currentCourse][currentNextCourse] = 1;
        }
      } else {
        counters[currentCourse] = {[currentNextCourse]:1};
      }
    }
  }
  const ans = {};
  Object.keys(counters).forEach(coursename => {
    ans[coursename] = Object.keys(counters[coursename]).reduce((a,b) => {
      const c1 = counters[coursename][a] || 0;
      const c2 = counters[coursename][b] || 0;
      if(c1 >= c2) {
        return a
      }
      return b
    }, "")
  })
  console.log(counters);
  console.log(ans);
}
solve_4();
