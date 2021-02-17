import React, { useState, useCallback, useEffect, useReducer } from "react";
import {
  Search,
  Button,
  Divider,
  Header,
  Grid,
  Card,
  Icon,
  Label
} from "semantic-ui-react";
import { Link } from "react-router-dom";
import _ from "lodash";

const initialState = {
  loading: false,
  results: [],
  value: "",
  resultsShown: []
};

function searchReducer(state, action) {
  switch (action.type) {
    case "CLEAN_QUERY":
      return initialState;
    case "START_SEARCH":
      return { ...state, loading: true, value: action.query };
    case "FINISH_SEARCH":
      return { ...state, loading: false, results: action.results };
    case "UPDATE_SELECTION":
      return { ...state, value: action.selection };
    case "UPDATE_SHOWN_RESULTS":
      return { ...state, resultsShown: action.data };
    default:
      throw new Error();
  }
}

const ContentCard = ({
  title,
  time,
  uploadedBY,
  approvedBY,
  tags,
  id,
  courseID,
  isSearched
}) => {
  let date = new Date(time * 1000).toLocaleString("en-US", {
    month: "long",
    year: "numeric"
  });

  console.log("isSearched:", isSearched, ", title:", title);

  return (
    <div>
      <Link to={"/course:" + courseID + "/content:" + id}>
        <Card
          fluid
          className="Content"
          style={{ backgroundColor: isSearched ? "#fffa63" : "" }}
        >
          <Card.Content>
            <Card.Header>{title}</Card.Header>
          </Card.Content>
          <Card.Content description>
            uploaded by <b>{uploadedBY.name}</b>
            {approvedBY ? "and approved by" : ""}{" "}
            <b>{approvedBY ? approvedBY.name : ""}</b> in time <b>{date}</b>
          </Card.Content>
          <Card.Content extra>
            {tags && tags.map(tag => {
              return (
                <Label style={{ marginBottom: 5 }}>
                  <Icon name="hashtag" /> {tag}
                </Label>
              );
            })}
          </Card.Content>
        </Card>
      </Link>
    </div>
  );
};

function ContentsPart({ contents, id }) {
  const [state, dispatch] = React.useReducer(searchReducer, initialState);
  const { loading, results, value, resultsShown } = state;

  if (resultsShown.length == 0) {
    dispatch({
      type: "UPDATE_SHOWN_RESULTS",
      data: contents
    });
  }

  const timeoutRef = React.useRef();
  const handleSearchChange = React.useCallback((e, data) => {
    clearTimeout(timeoutRef.current);
    dispatch({ type: "START_SEARCH", query: data.value });

    timeoutRef.current = setTimeout(() => {
      if (data.value.length === 0) {
        dispatch({ type: "CLEAN_QUERY" });
        return;
      }

      const re = new RegExp(_.escapeRegExp(data.value), "i");

      const isMatch = result => {
        var tag;
        console.log("tags:", result.tags);
        for (tag of result.tags) {
          let check = re.test(tag);
          console.log("checkkkk:", check, ", tag:", tag);
          if (check) {
            return true;
          }
        }
        return false;
      };

      let newArray = _.filter(contents, function(ss) {
        let check = isMatch(ss);
        return check;
      });

      dispatch({
        type: "FINISH_SEARCH",
        results: newArray
      });
    }, 300);
  }, []);
  React.useEffect(() => {
    return () => {
      clearTimeout(timeoutRef.current);
    };
  }, []);

  console.log("results shown:", resultsShown);
  console.log("results:", results);

  return (
    <div>
      <Grid columns={2} stackable>
        <Grid.Column floated="left">
          <Search
            showNoResults={false}
            loading={loading}
            onResultSelect={(e, data) => {
              dispatch({
                type: "UPDATE_SELECTION",
                selection: data.result.title
              });
            }}
            onSearchChange={handleSearchChange}
            results={[]}
            value={value}
          />
        </Grid.Column>
        {/* <Grid.Column floated="right">
          <Button
            onClick={() => {
              dispatch({ type: "UPDATE_SHOWN_RESULTS", data: results });
            }}
          >
            check
          </Button>
        </Grid.Column> */}
      </Grid>

      <Divider horizontal>
        <Header textAlign="left">
          <Icon name="video play" />
          Videos
        </Header>
      </Divider>
      <Grid columns={2} stackable>
        {resultsShown.map(content => {
          let isSearched = results.includes(content, c => {
            return c.id === content.id;
          });
          console.log("isSearched:", isSearched);
          return (
            <Grid.Column textAlign="left">
              <ContentCard
                key={content.id}
                title={content.title}
                time={content.timestamp}
                uploadedBY={content.uploadedBY}
                approvedBY={content.approvedBY}
                tags={content.tags}
                id={content.id}
                courseID={id}
                isSearched={isSearched}
              />
            </Grid.Column>
          );
        })}
      </Grid>
      <Divider horizontal>
        <Header textAlign="left">
          <Icon name="file" />
          Invetories
        </Header>
      </Divider>
      <Grid columns={1} textAlign="left">
        <Grid.Column>Sample inventory</Grid.Column>
      </Grid>
    </div>
  );
}

export default ContentsPart;
