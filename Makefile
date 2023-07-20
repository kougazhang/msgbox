commit: format test
	git add -A .
	git commit -m "$(info)"
	git push origin master
cur:
	git tag -l | sort -r|head -n1
tag:
	git tag $(tag)
	git push origin $(tag)
format:
	@find ./ -type f -name "*.go" -exec gofmt -w {} \;
test:
	go test ./...
init:
	go mod init github.com/kougazhang/$(project)
	echo '.idea' >> .gitignore
	echo "# $(project)" >> README.md
	git init
	git remote add origin git@github.com:kougazhang/$(project).git
