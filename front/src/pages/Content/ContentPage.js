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
  Segment,
  List
} from "semantic-ui-react";
import "./ContentPage.css";

<<<<<<< HEAD
// const STREAM_MUTATION = gql`
//   mutation Stream($vurl: String!) {
//     stream(vurl: $vurl) {

//     }
//   }
// `;
=======
const contentPageFrameLStyle={
  borderColor: "#0021a3",
  position: 'absolute',
  margin: "auto",
  top: "100px",
  left: "5%",
  width: "90%",
  height: "80%",
}

const contentPageSegmentLStyle={
  borderColor: "#0021a3",
  position: 'absolute',
  width: "100%",
  height: "100%",
  padding: "30px",
  overflow: 'auto',
  backgroundColor: '#1b1c1d'
}

const leftPanelFrameLStyle={
  height: "70vh",
}

const rightPanelFrameLStyle={
  height: "73vh",
  backgroundColor: "#ffffff",
  borderRadius:"0px"
}
>>>>>>> 498756b29f60fc1fa3c3a03086965a00ce3367f1

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

// createComment(username:String, contentID:String!, repliedAtID:String, target:TargetComment!): CreateCommentPayLoad!

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

const Comment = ({ comment, contentID }) => {
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
      console.log("localData in creating comment:", localData);

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
      console.log("data in cache--------", data);
      console.log("crreate comment:", createComment);
    },
    onCompleted: ({ createComment }) => {
      console.log("createCommenttttt:", createComment);
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
  console.log("contentID:", contentID, ", courseID:", courseID);

  const [newComment, setNewComment] = useState("");

  const { data, loading, error } = useQuery(CONTENT_QUERY, {
    variables: {
      id: contentID
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
      console.log("222 localData in creating comment:", localData);

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
      console.log("222data in cache--------", data);
      console.log("222crreate comment:", createComment);
    },
    onCompleted: ({ createComment }) => {
      console.log("222createCommenttttt:", createComment);
    }
  });

  console.log("data:", data);

  return (
    <div style={contentPageFrameLStyle}>
      {!loading && (
<<<<<<< HEAD
        <Segment style={{ top: 70, position: "absolute" }}>
=======
        <Segment raised style={contentPageSegmentLStyle}>
>>>>>>> 498756b29f60fc1fa3c3a03086965a00ce3367f1
          <Grid columns={2} textAlign="center" fluid stackable>
            <Grid.Column>
              <Segment inverted  style={leftPanelFrameLStyle}>
                <video width="100%" controls>
                  <source
                    src={
                      // "https://s70.upera.net/2751313-0-WonderWoman4849193-480.mp4?owner=2640789&ref=1794068"
                      data.content.vurl
                    }
                    type="video/mp4"
                  />
                  {/* <source  type="" /> */}
                  Your browser does not support HTML video.
                </video>
<<<<<<< HEAD
                <Container textAlign="left">
                  <List horizontal>
                    <List.Item>
                      <a href={data.content.vurl} download>
                        <Button color="blue" icon floated="right">
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
            <Grid.Column style={{ height: "100%", overflow: "auto" }}>
              <Segment>
=======
                <Container text textAlign="left" style={{overflow:'auto', paddingTop:"20px"}}>
                  <h1>{data.content.title}</h1>
                  <p>
                    {data.content.description}
                  </p>
                </Container>
              </Segment>
            </Grid.Column>
            <Grid.Column style={{height: "100%", overflow: "auto"}}>
              <Segment style={rightPanelFrameLStyle}>
>>>>>>> 498756b29f60fc1fa3c3a03086965a00ce3367f1
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
                          <Comment comment={comment} contentID={contentID} />
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
