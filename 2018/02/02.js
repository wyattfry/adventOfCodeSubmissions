const fs = require('fs');

const charCounts = (str) => {
  const counts = {};
  [...str].forEach((l) => {
    if (counts.hasOwnProperty(l)) {
      counts[l]++;
    } else {
      counts[l] = 1;
    }
  });
  return counts;
}

const containsN = (str, n) => {
  const counts = charCounts(str);
  for (let c of Object.values(counts)) {
    if (c === n) {
      return true;
    }
  }
  return false;
}

const countDiffs= (a, b) => {
  let diffs = 0;
  // assume a and b same length
  [...a].forEach((l, i) => {
    diffs += l !== [...b][i];
  });
  return diffs;
}


fs.readFile('input', 'utf8', (err, data) => {
  if (err) {
    console.error(err);
    process.exit(1);
  }
  const ids = data.split('\n');
  let count2 = 0;
  let count3 = 0;
  ids.forEach((id) => {
    count2 += containsN(id, 2) ? 1 : 0;
    count3 += containsN(id, 3) ? 1 : 0;
  });
  console.log(count2, count3);


  // Part 2
  //
  for (let a = 0; a < ids.length - 1; a++) {
    for (let b = a + 1; b < ids.length; b++) {
      if (countDiffs(ids[a], ids[b]) == 1) {
        console.log(ids[a] + '\n' + ids[b]);
        process.exit();
      }
    }
  }
});



