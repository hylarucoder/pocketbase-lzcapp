rm -r ./dist && mkdir -p dist
cp run.sh ./dist/run.sh
GOOS=linux GOARCH=amd64 go build -o ./dist/main
chmod +x ./dist/main
# cd ../ui && npx vite build --outDir ../dist/web
