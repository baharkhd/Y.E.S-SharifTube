import React, { useState } from "react";
import { Grid, Form, Segment, Message, Input, Button } from "semantic-ui-react";
import gql from "graphql-tag";
import { useMutation } from "@apollo/client";
import { useHistory } from "react-router-dom";
import constants from "../../constants.js";

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

const SIGNUP_REGISTER = gql`
  mutation CreateUser(
    $username: String!
    $password: String!
    $name: String
    $email: String
  ) {
    createUser(
      target: {
        username: $username
        password: $password
        name: $name
        email: $email
      }
    ) {
      __typename
      ... on User {
        username
      }
      ... on Exception {
        message
      }
    }
  }
`;

const USERS_QUERY = gql`
  {
    users(start: 0, amount: 100) {
      username
    }
  }
`;

const RegisterForm = props => {
  const [state, setState] = useState({
    name: "",
    username: "",
    email: "",
    password: "",
    confirmPass: "",
    error: ""
  });

  const history = useHistory();

  async function changeToken(token) {
    await props.setToken(token);
    history.push("/dashboard");
  }

  const [login] = useMutation(LOGIN_MUTATION, {
    variables: {
      username: state.username,
      password: state.password
    },
    update(cache, { data: { login } }) {
      // console.log("update in login:", login);
      // console.log("cache in login update fuunction:", cache);
    },
    onCompleted: ({ login }) => {
      if (login.__typename == "Token") {
        props.setUsername(state.username);
        // props.setToken(login.token);
        // history.push("/dashboard");
        props.makeNotif("Success", "You successfully loged in .", "success");
        changeToken(login.token);
      } else {
        switch (login.__typename) {
          case "UserPassMissMatchException":
            // setState({ ...state, error: constants.USER_PASS_MISMATCH });
            props.makeNotif("Error", constants.USER_PASS_MISMATCH, "danger");
            break;
          case "InternalServerException":
            // alert(constants.INTERNAL_SERVER_EXCEPTION);
            props.makeNotif("Error", "Login was not successfull .", "danger");
            // setState({ ...state, error: constants.INTERNAL_SERVER_EXCEPTION });
            break;
        }
      }
    }
  });

  const [createUser] = useMutation(SIGNUP_REGISTER, {
    variables: {
      username: state.username,
      name: state.name,
      email: state.email,
      password: state.password
    },
    onCompleted: ({ createUser }) => {
      console.log("createUser:", createUser);
      if (createUser.__typename == "User") {
        login();
      } else {
        if (createUser.__typename === "DuplicateUsernameException") {
          props.makeNotif(
            "Error",
            constants.DUPLICATE_USERNAME_ERROR,
            "danger"
          );
        } else {
          props.makeNotif("Error", "Signup was not successfull .", "danger");
        }
      }
    }
  });

  return (
    <Form>
      <Segment>
        <Form.Input
          icon="smile"
          iconPosition="left"
          placeholder="Enter your name"
          control={Input}
          onChange={e => {
            setState({
              ...state,
              name: e.target.value
            });
          }}
        />
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
        />
        <Form.Input
          icon="mail"
          iconPosition="left"
          placeholder="Enter your email"
          control={Input}
          onChange={e => {
            setState({
              ...state,
              email: e.target.value
            });
          }}
        />
        <Form.Input
          icon="lock"
          iconPosition="left"
          type="password"
          placeholder="Choose a password"
          control={Input}
          onChange={e => {
            setState({
              ...state,
              password: e.target.value
            });
          }}
        />
        <Form.Input
          icon="lock"
          iconPosition="left"
          type="password"
          placeholder="Repeat your password"
          control={Input}
          onChange={e => {
            setState({
              ...state,
              confirmPass: e.target.value
            });
          }}
        />
        <Form.Button
          fluid
          color="blue"
          content="Register"
          control={Button}
          onClick={() => {
            //   handleRegister();
            if (
              state.name.trim() !== "" &&
              state.username.trim() !== "" &&
              state.email.trim() !== "" &&
              state.password.trim() !== "" &&
              state.confirmPass.trim() !== ""
            ) {
              if (state.password === state.confirmPass) {
                createUser();
              } else {
                props.makeNotif("Error", constants.PASSWORDS_DIFFER, "danger");
              }
            } else {
              props.makeNotif("Error!", constants.EMPTY_FIELDS, "danger");
            }
          }}
        />
      </Segment>
      <Message>
        Already have an account? <a href="/login">Login</a>
      </Message>
    </Form>
  );
};

function Signup(props) {
  return (
    <div style={{ top: "80px", position: "absolute", width: "100%" }}>
      <Grid
        centered
        verticalAlign="middle"
        textAlign="center"
        style={{ height: "100vh" }}
      >
        <Grid.Row>
          <Grid.Column
            style={{ maxWidth: 450, marginRight: 20, marginLeft: 20 }}
          >
            <RegisterForm
              setUsername={props.setUsername}
              setToken={props.setToken}
              makeNotif={props.makeNotif}
            />
          </Grid.Column>
        </Grid.Row>
      </Grid>
    </div>
  );
}

export default Signup;
