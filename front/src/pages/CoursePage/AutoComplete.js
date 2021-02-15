import React, { Component, useState } from "react";
import PropTypes from "prop-types";
import { Segment, Input, List, Button, Icon } from "semantic-ui-react";
import { gql, useMutation } from "@apollo/client";

const ADD_TA_MUTATION = gql`
  mutation AddTA($courseID: String!, $targetUsername: String!) {
    promoteUserToTA(courseID: $courseID, targetUsername: $targetUsername) {
      __typename
      ... on Course {
        id
        title
        summary
        createdAt
      }
      ... on Exception {
        message
      }
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

  const [promoteUserToTA] = useMutation(ADD_TA_MUTATION, {
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
