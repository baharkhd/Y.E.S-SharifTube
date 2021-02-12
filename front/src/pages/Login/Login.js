import React, { Component, useState } from "react";
import { Grid, Form, Segment, Message, Input, Button } from "semantic-ui-react";
import gql from "graphql-tag";
import { useMutation } from "@apollo/client";
import { useHistory } from 'react-router-dom'

const LOGIN_MUTATION = gql`
  mutation Login($username: String!, $password: String!) {
    login(input: { username: $username, password: $password }) {
      __typename
      ... on Token {
        token
      }
    }
  }
`;

const LoginForm = props => {
  const [state, setState] = useState({
    username: " ",
    password: " ",
    error: ""
  });

  const history = useHistory()

  const [login] = useMutation(LOGIN_MUTATION, {
    variables: {
      username: state.username,
      password: state.password
    },
    onCompleted: ({ login }) => {
      console.log("login response:", login);
    }
  });

  function handleLogin() {
    if (state.username && state.password) {
      console.log("handliing login?????????");
      login();
      setState({ ...state, error: "" });
      history.push("/dashboard")
    }
  }

  return (
    <div>
      <Form>
        <Segment>
          <Form.Input
            icon="user"
            iconPosition="left"
            placeholder="Enter your username"
            control={Input}
            onChange={e => {
              setState({
                ...state,
                username: e.target.value
              });
            }}
            error={
              !state.username && {
                content: "Please enter a username!",
                pointing: "below"
              }
            }
          />
          <Form.Input
            icon="lock"
            iconPosition="left"
            placeholder="Enter your password"
            control={Input}
            type="password"
            onChange={e => {
              setState({
                ...state,
                password: e.target.value
              });
            }}
            error={
              !state.password && {
                content: "Please enter a password please!",
                pointing: "below"
              }
            }
          />
          <Form.Button
            fluid
            color="blue"
            content="Login"
            control={Button}
            onClick={() => {
              handleLogin();
            }}
          />
        </Segment>
      </Form>
      <Message>
        New to us? <a href="/signup">Sign Up</a>
      </Message>
      {state.error !== "" && <Message negative>{state.error}</Message>}
    </div>
  );
};

function Login() {
  return (
    <div style={{ top: "80px", position: "absolute", width: "100%" }}>
      <Grid
        centered
        style={{ height: "100vh" }}
        verticalAlign="middle"
        textAlign="center"
      >
        <Grid.Row>
          <Grid.Column
            style={{
              maxWidth: 450,
              marginRight: 20,
              marginLeft: 20
            }}
          >
            <LoginForm />
          </Grid.Column>
        </Grid.Row>
      </Grid>
    </div>
  );
}

export default Login;
