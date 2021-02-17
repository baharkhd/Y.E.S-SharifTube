import React, { useState } from "react";
import {
  Embed,
  Placeholder,
  Segment,
  Grid,
  Feed,
  Icon,
  Divider,
  Container,
  Input,
  Button
} from "semantic-ui-react";
import "./ContentPage.css";
import { useParams } from "react-router-dom";
import { useQuery, gql, useMutation } from "@apollo/client";
import _ from "lodash";
import poster from "./test2.png";

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
    year: "numeric",
    day: "numeric"
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
    year: "numeric",
    day: "numeric"
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

// mutation CreateComment(
//   $contentID: String!
//   $repliedAtID: String
//   $body: String!
// ) {
//   createComment(
//     contentID: $contentID
//     repliedAtID: $repliedAtID
//     body: $body
//   )

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
    <div>
      {!loading && data && (
        <Segment style={{ top: 70, overflow: "hidden", borderRadius: 0 }}>
          <Grid columns={2} textAlign="center" fluid stackable>
            <Grid.Column>
              <Segment>
                <Container textAlign="center" fluid>
                  <video width="100%" controls poster={poster}>
                    <source
                      src={
                        data.content.vurl
                        // "https://s70.upera.net/2751313-0-WonderWoman4849193-480.mp4?owner=2640789&ref=1794068"
                      }
                      type="video/mp4"
                    />
                  </video>
                  <Container textAlign="left" style={{ left: 20 }}>
                    <h1>{data.content.title}</h1>
                    <p>
                      {data.content.description}
                      {/* Lorem ipsum dolor sit amet, consectetuer adipiscing elit.
                    Aenean commodo ligula eget dolor. Aenean massa strong. Cum
                    sociis natoque penatibus et magnis dis parturient montes,
                    nascetur ridiculus mus. Donec quam felis, ultricies nec,
                    pellentesque eu, pretium quis, sem. Nulla consequat massa
                    quis enim. Donec pede justo, fringilla vel, aliquet nec,
                    vulputate eget, arcu. In enim justo, rhoncus ut, imperdiet
                    a, venenatis vitae, justo. Nullam dictum felis eu pede link
                    mollis pretium. Integer tincidunt. Cras dapibus. Vivamus
                    elementum semper nisi. Aenean vulputate eleifend tellus.
                    Aenean leo ligula, porttitor eu, consequat vitae, eleifend
                    ac, enim. Aliquam lorem ante, dapibus in, viverra quis,
                    feugiat a, tellus. Phasellus viverra nulla ut metus varius
                    laoreet. Quisque rutrum. Aenean imperdiet. Etiam ultricies
                    nisi vel augue. Curabitur ullamcorper ultricies nisi. */}
                    </p>
                  </Container>
                </Container>
              </Segment>
            </Grid.Column>
            <Grid.Column>
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
