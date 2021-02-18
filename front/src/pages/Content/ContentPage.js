import { gql, useMutation, useQuery } from "@apollo/client";
import _ from "lodash";
import React, { useState } from "react";
import { useParams } from "react-router-dom";
import {
  Button,
  Container,
  Feed,
  Grid,
  Icon,
  Input,
  List,
  Segment
} from "semantic-ui-react";

const contentPageFrameLStyle = {
  borderColor: "#0021a3",
  position: "absolute",
  margin: "auto",
  top: "100px",
  left: "5%",
  width: "90%",
  height: "80%"
};

const contentPageSegmentLStyle = {
  borderColor: "#0021a3",
  position: "absolute",
  width: "100%",
  height: "100%",
  padding: "30px",
  overflow: "auto",
  backgroundColor: "#1b1c1d"
};

const leftPanelFrameLStyle = {
  height: "70vh"
};

const rightPanelFrameLStyle = {
  height: "73vh",
  backgroundColor: "#ffffff",
  borderRadius: "0px"
};

const CONTENT_QUERY = gql`
  query GetContent($id: String!) {
    content(id: $id) {
      id
      title
      description
      vurl
      tags
      timestamp
      comments {
        id
        author {
          id
          name
          username
        }
        body
        timestamp
        contentID
        replies {
          # id
          author {
            id
            name
            username
          }
          body
          timestamp
        }
      }
    }
  }
`;

const ADD_COMMENT_MUTATION = gql`
  mutation CreateComment(
    $contentID: String!
    $repliedAtID: String
    $body: String!
  ) {
    createComment(
      contentID: $contentID
      repliedAtID: $repliedAtID
      target: { body: $body }
    ) {
      __typename
      ... on Comment {
        author {
          id
          name
          username
        }
        id
        body
        timestamp
        contentID
        replies {
          author {
            id
            name
            username
          }
          body
          timestamp
        }
      }
      ... on Reply {
        # id
        body
        timestamp
        author {
          id
          name
          username
        }
        commentID
      }
      ... on Exception {
        message
      }
    }
  }
`;

const Reply = ({ author, body, time }) => {
  let date = new Date(time * 1000).toLocaleString("en-US", {
    month: "long",
    year: "numeric"
  });
  return (
    <Feed.Event>
      <Feed.Content>
        <Feed.Summary>
          <Icon name="reply" />
          <Feed.User>{author}</Feed.User> {body}
          <Feed.Date>{date}</Feed.Date>
        </Feed.Summary>
      </Feed.Content>
    </Feed.Event>
  );
};

const Comment = ({ comment, contentID, makeNotif }) => {
  let date = new Date(comment.timestamp * 1000).toLocaleString("en-US", {
    month: "long",
    year: "numeric"
  });

  const [reply, setReply] = useState("");

  // For replies
  const [createComment] = useMutation(ADD_COMMENT_MUTATION, {
    update: (cache, { data: { createComment } }) => {
      const data = cache.readQuery({
        query: CONTENT_QUERY,
        variables: {
          id: contentID
        }
      });

      let localData = _.cloneDeep(data);

      localData.content.comments = localData.content.comments.map(cm => {
        return cm.id === createComment.commentID
          ? {
              ...cm,
              replies: [...(cm.replies ? cm.replies : []), createComment]
            }
          : cm;
      });

      cache.writeQuery({
        query: CONTENT_QUERY,
        data: {
          ...localData
        }
      });
    },
    onCompleted: ({ createComment }) => {
      if (createComment.__typename === "Reply") {
        // Successfull
      } else {
        makeNotif("Error!", createComment.message, "danger");
      }
      // console.log("createCommenttttt:", createComment);
    }
  });

  return (
    <Feed.Event>
      <Feed.Content>
        <Feed.Summary>
          <Feed.User>{comment.author.name}</Feed.User> {comment.body}
          <Feed.Date>{date}</Feed.Date>
        </Feed.Summary>
        <Feed.Meta>
          {comment.replies &&
            comment.replies.map(reply => {
              return (
                <Reply
                  author={reply.author.name}
                  time={reply.timestamp}
                  body={reply.body}
                />
              );
            })}
          <Feed.Event>
            <Feed.Content>
              <Feed.Summary>
                <Input
                  action={
                    <Button
                      color="black"
                      icon="comment"
                      onClick={() => {
                        if (reply != "") {
                          createComment({
                            variables: {
                              contentID: contentID,
                              body: reply,
                              repliedAtID: comment.id
                            }
                          });
                          setReply("");
                        }
                      }}
                    />
                  }
                  placeholder="Add a reply ..."
                  onChange={e => {
                    setReply(e.target.value);
                  }}
                  value={reply}
                />
              </Feed.Summary>
            </Feed.Content>
          </Feed.Event>
        </Feed.Meta>
      </Feed.Content>
    </Feed.Event>
  );
};

function ContentPage(props) {
  let { courseID, contentID } = useParams();
  courseID = courseID.substring(1);
  contentID = contentID.substring(1);

  const [newComment, setNewComment] = useState("");

  const { data, loading, error } = useQuery(CONTENT_QUERY, {
    variables: {
      id: contentID
    },
    onCompleted({ content }) {
      if (content.__typename === "Content") {
      }
    }
  });

  // For main comments
  const [createComment] = useMutation(ADD_COMMENT_MUTATION, {
    variables: {
      contentID: contentID,
      body: newComment
    },
    update: (cache, { data: { createComment } }) => {
      const data = cache.readQuery({
        query: CONTENT_QUERY,
        variables: {
          id: contentID
        }
      });

      let localData = _.cloneDeep(data);

      localData.content.comments = [
        ...(localData.content.comments ? localData.content.comments : []),
        createComment
      ];

      cache.writeQuery({
        query: CONTENT_QUERY,
        data: {
          ...localData
        }
      });
    },
    onCompleted: ({ createComment }) => {
      if (createComment.__typename === "Comment") {
        // Successfull
      } else {
        props.makeNotif("Error!", createComment.message, "danger");
      }
    }
  });

  return (
    <div style={contentPageFrameLStyle}>
      {!loading && (
        <Segment raised style={contentPageSegmentLStyle}>
          <Grid columns={2} textAlign="center" fluid stackable>
            <Grid.Column>
              <Segment inverted style={leftPanelFrameLStyle}>
                <video width="100%" controls>
                  <source src={data.content.vurl} type="video/mp4" />
                  Your browser does not support HTML video.
                </video>
                <Container textAlign="left">
                  <List horizontal>
                    <List.Item>
                      <a href={data.content.vurl} download>
                        <Button
                          color="blue"
                          icon
                          floated="right"
                          loading={loading}
                        >
                          <Icon name="download" />
                        </Button>
                      </a>
                    </List.Item>

                    <List.Item>
                      <h1>{data.content.title}</h1>
                    </List.Item>
                  </List>

                  <p>{data.content.description}</p>
                </Container>
              </Segment>
            </Grid.Column>
            <Grid.Column style={{ height: "100%", overflow: "auto", top: 13 }}>
              <Segment>
                <Input
                  fluid
                  type="string"
                  action={
                    <Button
                      color="blue"
                      icon="plus"
                      onClick={() => {
                        if (newComment) {
                          createComment();
                          setNewComment("");
                        }
                      }}
                    />
                  }
                  onChange={e => {
                    setNewComment(e.target.value);
                  }}
                  actionPosition="right"
                  placeholder="Add a comment ..."
                  value={newComment}
                />
                <Feed>
                  {data.content.comments &&
                    data.content.comments.map(comment => {
                      return (
                        <>
                          <Comment
                            comment={comment}
                            contentID={contentID}
                            makeNotif={props.makeNotif}
                          />
                          <p></p>
                        </>
                      );
                    })}
                </Feed>
              </Segment>
            </Grid.Column>
          </Grid>
        </Segment>
      )}
    </div>
  );
}

export default ContentPage;
