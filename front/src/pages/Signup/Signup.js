import React, {useState} from "react";
import {Grid, Form, Segment, Message, Input, Button} from "semantic-ui-react";
import gql from "graphql-tag";
import {useMutation} from "@apollo/client";
import {useHistory} from "react-router-dom";
import constants from "../../constants.js";


const signUpBodyLStyle = {
    backgroundColor: '#abe0fd',
    height: '100vh'
}

const signUpFormContainerStyle = {
    top: "80px",
    position: "absolute",
    width: "100%"
}

const signUpFormInputLStyle = {
    color: '#023849',
}

const signUpFormLStyle = {
    backgroundColor: '#023849',
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
        update(cache, {data: {login}}) {
            // console.log("update in login:", login);
            // console.log("cache in login update fuunction:", cache);
        },
        onCompleted: ({login}) => {
            if (login.__typename == "Token") {
                console.log("login:", login);
                console.log("token in logiin:", login.token);
                props.setUsername(state.username);
                // props.setToken(login.token);
                // history.push("/dashboard");
                changeToken(login.token);
            } else {
                switch (login.__typename) {
                    case "UserPassMissMatchException":
                        alert(constants.USER_PASS_MISMATCH);
                        setState({...state, error: constants.USER_PASS_MISMATCH});
                        break;
                    case "InternalServerException":
                        alert(constants.INTERNAL_SERVER_EXCEPTION);
                        setState({...state, error: constants.INTERNAL_SERVER_EXCEPTION});
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
        onCompleted: ({createUser}) => {
            console.log("createUser:", createUser);
            if (createUser.__typename == "User") {
                login();
            } else {
                alert(createUser.__typename);
            }
        }
    });

    return (
        <Form>
            <Segment style={signUpFormLStyle}>
                <Form.Input style={signUpFormInputLStyle}
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
                <Form.Input style={signUpFormInputLStyle}
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
                <Form.Input style={signUpFormInputLStyle}
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
                <Form.Input style={signUpFormInputLStyle}
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
                <Form.Input style={signUpFormInputLStyle}
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
                        if (state.password === state.confirmPass) {
                            createUser();
                        } else {
                            alert("Passwords mismatch")
                        }

                    }}
                />
            </Segment>
            <Message>
                Already have an account? <a href="/login">Login</a>
            </Message>
            {state.error !== "" && <Message negative>{state.error}</Message>}
        </Form>
    );
};

function Signup(props) {
    return (
        <div style={signUpBodyLStyle}>
            <div style={signUpFormContainerStyle}>
                <Grid
                    centered
                    verticalAlign="middle"
                    textAlign="center"
                    style={{height: "70vh"}}
                >
                    <Grid.Row>
                        <Grid.Column
                            style={{maxWidth: 450, marginRight: 20, marginLeft: 20}}
                        >
                            <RegisterForm setUsername={props.setUsername} setToken={props.setToken}/>
                        </Grid.Column>
                    </Grid.Row>
                </Grid>
            </div>
        </div>
    );
}

export default Signup;
