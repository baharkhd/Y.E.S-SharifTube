import { gql, useMutation } from "@apollo/client";
import _ from 'lodash';
import React, { useState } from "react";
import { useParams } from "react-router-dom";
import { Button, Input, List } from "semantic-ui-react";

const ADD_TA_MUTATION = gql`
  mutation AddTA($courseID: String!, $targetUsername: String!) {
    promoteUserToTA(courseID: $courseID, targetUsername: $targetUsername) {
      __typename
      ... on Course {
        id
        title
        summary
        createdAt
        tas {
          username
          name
        }
      }
      ... on Exception {
        message
      }
    }
  }
`;

const COURSE_QUERY = gql`
  query GetCoursesByID($ids: [String!]!) {
    courses(ids: $ids) {
      id
      #   title
      #   summary
      #   contents {
      #     id
      #     title
      #     description
      #   }
      #   prof {
      #     name
      #     username
      #     email
      #   }
      tas {
        name
        username
      }
      #   students {
      #     username
      #   }
    }
  }
`;

function Autocomplete(props) {
  const [state, setState] = useState({
    activeSuggestion: 0,
    filteredSuggestions: [],
    showSuggestions: false,
    userInput: "",
    TA_username: ""
  });

  let { id } = useParams();
  id = id.substring(1);

  const [promoteUserToTA] = useMutation(ADD_TA_MUTATION, {
    update(cache, { data: { createPost } }) {
      const data = cache.readQuery({
        query: COURSE_QUERY,
        variables: {
          ids: [id]
        }
      });

      const localCourse = _.cloneDeep(data);
      console.log("localData in auto complete:", localCourse);
      console.log("newData:", {
        ...localCourse
      });

      cache.writeQuery({
        query: COURSE_QUERY,
        data: {
          ...localCourse
        }
      });
    },
    onCompleted: ({ promoteUserToTA }) => {
      console.log("promoteUserToTA------- :", promoteUserToTA);
    }
  });

  function onChange(e) {
    const { suggestions } = props;
    const userInput = e.currentTarget.value;

    const filteredSuggestions = suggestions.filter(
      suggestion =>
        suggestion.toLowerCase().indexOf(userInput.toLowerCase()) > -1
    );

    setState({
      activeSuggestion: 0,
      filteredSuggestions,
      showSuggestions: true,
      userInput: e.currentTarget.value
    });
  }

  function onClick(e) {
    setState({
      activeSuggestion: 0,
      filteredSuggestions: [],
      showSuggestions: false,
      userInput: e.currentTarget.innerText
    });
  }

  function onKeyDown(e) {
    const { activeSuggestion, filteredSuggestions } = state;

    if (e.keyCode === 13) {
      this.setState({
        activeSuggestion: 0,
        showSuggestions: false,
        userInput: filteredSuggestions[activeSuggestion]
      });
    } else if (e.keyCode === 38) {
      if (activeSuggestion === 0) {
        return;
      }

      setState({ activeSuggestion: activeSuggestion - 1 });
    } else if (e.keyCode === 40) {
      if (activeSuggestion - 1 === filteredSuggestions.length) {
        return;
      }

      setState({ activeSuggestion: activeSuggestion + 1 });
    }
  }

  let suggestionsListComponent;
  if (state.showSuggestions && state.userInput) {
    if (state.filteredSuggestions.length) {
      suggestionsListComponent = (
        <ul class="suggestions">
          {state.filteredSuggestions.map((suggestion, index) => {
            let className;

            if (index === state.activeSuggestion) {
              className = "";
            }

            return (
              <List divided verticalAlign="middle">
                <List.Item>
                  <List.Content floated="right">
                    <Button
                      positive
                      onClick={() => {
                        promoteUserToTA({
                          variables: {
                            courseID: props.courseID,
                            targetUsername: suggestion
                          }
                        });
                      }}
                    >
                      Add
                    </Button>
                  </List.Content>
                  <List.Content>{suggestion}</List.Content>
                </List.Item>
              </List>
            );
          })}
        </ul>
      );
    } else {
      suggestionsListComponent = (
        <div class="no-suggestions">
          <em>No suggestions</em>
        </div>
      );
    }
  }

  return (
    <React.Fragment>
      <Input
        type="search"
        onChange={onChange}
        onKeyDown={onKeyDown}
        value={state.userInput}
        placeholder="Search usernames"
      />
      {suggestionsListComponent}
    </React.Fragment>
  );
}

export default Autocomplete;
