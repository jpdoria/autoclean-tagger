# About

Add "autoclean=true" tag to all objects in an S3 bucket.

# Requirements

- [Go](https://golang.org/dl/)
- Node.js: download from their [official website](https://nodejs.org/en/download/) or use [nvm](https://github.com/nvm-sh/nvm#installation-and-update).
- [Serverless Framework](https://github.com/serverless/serverless/)
- Make sure you have a valid AWS account and your [AWS credential file](https://aws.amazon.com/blogs/security/a-new-and-standardized-way-to-manage-credentials-in-the-aws-sdks/) is properly installed.

# Go Setup

Please follow the instructions [here](https://golang.org/doc/install).

# Cloning Repository

```bash
git clone git@github.com:jpdoria/autoclean-tagger.git $HOME/go/src/autoclean-tagger
```

# Installing Dependencies

```bash
cd $HOME/go/src/autoclean-tagger
go get -v
```

# Serverless Setup

```bash
npm install -g serverless
```

# Deployment

```
make deploy
```
