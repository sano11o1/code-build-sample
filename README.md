# code-build-sample
- codebuildのプロジェクトを手動で作成する✅
- 手動でbuildを開始する✅
- リポジトリに保存したbuildspecからgo buildしlambda関数を更新する✅
  - buildspec.ymlから`aws lambda update-function-code`を実行するにはcodebuildのポリシーにAWSLambda_FullAccessを追加する必要あり
- mainブランチにPRをマージした時に自動でbuildが走るようにする✅
  - codebuildプロジェクト -> ソースを編集でリポジトリの選択をGitHubアカウントのリポジトリにする
  - ソースバージョンを `red/head/main`にする
  - ウェブフックのチェックボックスをオンにし、イベントタイプにプッシュを指定、ビルド条件に`HEAD_REF: ^refs/heads/main$`を指定
