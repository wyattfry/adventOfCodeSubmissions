const fs = require('fs');
/*
 which guard is asleep for the most minutes
 which minute they are most often asleep

[1518-05-12 00:46] wakes up

1. sort the observations chronologically
2. total # minutes asleep for each guard
3. find guard with most minutes asleep
4. find which minute they are most often asleep

*/

//const file = './test_input';
const file = './input';

const getGuardsSleepiestMinAndCount = (lines, guard) => {
    // returns [sleepiestMinute, count]
    // e.g. [10, 22] = guard was asleep 22 times on the 10th minute
    const minutes = new Array(60);
    minutes.fill(0);
    let sleepiestMinute = -1;
    let sleepiestMinuteCount = -1;
    lines.forEach((line) => {
        // [1518-05-12 00:46] wakes up
        if (line.includes('begins shift')) {
            guardNumber = line.split(' ')[3].substring(1);
        }
        if (line.includes('falls asleep') && guardNumber == guard) {
            sleepTime = parseInt(line.split(' ')[1].substring(3,5)); 
        }
        if (line.includes('wakes up') && guardNumber == guard) {
            awakeTime = parseInt(line.split(' ')[1].substring(3,5));
            for (let m = sleepTime; m < awakeTime; m++) {
                minutes[m]++;
                if (minutes[m] > sleepiestMinuteCount) {
                    sleepiestMinute = m;
                    sleepiestMinuteCount = minutes[m];
                }
            }            
        }
    });
    return [sleepiestMinute, sleepiestMinuteCount];
}

fs.readFile(file, 'utf8', (error, data) => {
    if (error) {
        console.error(error);
        process.exit(1);
    }
    const lines = data.split('\n');
    // sort chrono
    lines.sort();   
    const guardSleepLog = {}
    let guardNumber;
    let sleepTime;
    let awakeTime;
    let minsSlept;
    let mostMinsSlept = 0;
    let sleepiestGuard;
    let sleepLog = [];
    lines.forEach((line) => {
        // [1518-05-12 00:46] wakes up
        if (line.includes('begins shift')) {
            guardNumber = line.split(' ')[3].substring(1);
        }
        if (line.includes('falls asleep')) {
            sleepTime = parseInt(line.split(' ')[1].substring(3,5)); 
        }
        if (line.includes('wakes up')) {
            awakeTime = parseInt(line.split(' ')[1].substring(3,5));
            minsSlept = awakeTime - sleepTime;
            sleepLog.push(minsSlept);
            guardSleepLog[guardNumber] = (guardSleepLog[guardNumber] + minsSlept) || minsSlept;

            if (guardSleepLog[guardNumber] > mostMinsSlept) {
                mostMinsSlept = guardSleepLog[guardNumber];
                sleepiestGuard = guardNumber;
            }
        }
    });
    console.log('guard:', sleepiestGuard);

    // Now that we know the guard that slept the most minutes total
    // Find out which minute were they most often asleep across all
    // their shifts. I imagine an array for each minute between
    // 00:00 and 00:59, could just be an array with 60 elements

    const res = getGuardsSleepiestMinAndCount(lines, sleepiestGuard);
    console.log(JSON.stringify(res));

    /*
--- Part Two ---
Strategy 2: Of all guards, which guard is most frequently asleep on the same minute?
In the example above, Guard #99 spent minute 45 asleep more than any other guard or minute - three times in total. (In all other cases, any guard spent any minute asleep at most twice.)
What is the ID of the guard you chose multiplied by the minute you chose? (In the above example, the answer would be 99 * 45 = 4455.)
    */
   let guard;
   let minute;
   let count = 0;
   let minAndCount = new Array(2);
   Object.keys(guardSleepLog).forEach(g => {
       minAndCount = getGuardsSleepiestMinAndCount(lines, g);
       if (minAndCount[1] > count) {
           minute = minAndCount[0];
           count = minAndCount[1];
           guard = g;
       }
   });
   console.log(guard, minute, guard * minute);
});
