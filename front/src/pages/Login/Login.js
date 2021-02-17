import React, {Component, useState} from "react";
import {Grid, Form, Segment, Message, Input, Button} from "semantic-ui-react";
import gql from "graphql-tag";
import {useMutation} from "@apollo/client";
import {useHistory} from "react-router-dom";
import constants from "../../constants.js";

const signInBodyLStyle = {
    backgroundColor: '#abe0fd',
    height: '100vh',
}

const signInFormContainerStyle = {
    top: "70px",
    position: "absolute",
    width: "100%"
}

const signInFormLStyle = {
    backgroundColor: '#023849',
}

const signInFormInputLStyle = {
    color: '#023849',
}

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
        update(cache, {data: {login}}) {
            console.log("update in login:", login)
            console.log("cache in login update fuunction:", cache)
        },
        onCompleted: ({login}) => {
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

    function handleLogin() {
        if (state.username && state.password) {
            // console.log("handliing login?????????");
            login();
            // setState({ ...state, error: "" });
            // history.push("/dashboard");
        }
    }
  function handleLogin() {
    if (state.username.trim() !== "" && state.password.trim() !== "") {
      login();
    } else {
      props.makeNotif("Error!", constants.EMPTY_FIELDS, "danger")
    }
  }

    return (
        <div>
            <Form>
                <Segment style={signInFormLStyle}>
                    <Form.Input style={signInFormInputLStyle}
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
                    <Form.Input style={signInFormInputLStyle}
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
           </div>
    );
};

function Login(props) {
    return (
        <div style={signInBodyLStyle}>
            <div style={signInFormContainerStyle}>
                <Grid
                    centered
                    style={{height: "60vh"}}
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
                          <LoginForm
                              setToken={props.setToken}
                              setUsername={props.setUsername}
                              makeNotif={props.makeNotif}
                          />
                        </Grid.Column>
                    </Grid.Row>
                </Grid>
            </div>
        </div>
    );
}

export default Login;
