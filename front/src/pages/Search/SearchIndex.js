import React from "react";
import SearchCourse from "./SearchCourse.js";
import faker from "faker";
import _ from "lodash";

// const source = _.times(5, () => ({
//     title: faker.company.companyName(),
//     description: faker.company.catchPhrase(),
//     image: faker.internet.avatar(),
//     price: faker.finance.amount(0, 100, 2, '$'),
//   }))

const source = [
  {
    title: "class1",
    summary: "summary1",
    id: "ID1"
  },
  {
    title: "class2",
    summary: "summary2",
    id: "ID2"
  },
  {
    title: "class3",
    summary: "summary3",
    id: "ID3"
  },
  {
    title: "class4",
    summary: "summary4",
    id: "ID4"
  },
  {
    title: "class5",
    summary: "summary5",
    id: "ID5"
  },
  {
    title: "class6",
    summary: "summary6",
    id: "ID6"
  },
  {
    title: "class7",
    summary: "summary7",
    id: "ID7"
  },
  {
    title: "class8",
    summary: "summary8",
    id: "ID8"
  }
];

function SearchIndex() {
  return <SearchCourse source={source} />;
}

export default SearchIndex;
