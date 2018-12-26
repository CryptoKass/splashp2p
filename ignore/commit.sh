#!/bin/bash
echo "[COMMIT WIZARD]"
echo -e -n "Commit Message:"
read CommitMessage


# update readme.md
echo "[TASK] Copying readme template to Project folders"

echo -ne "[----]"\\r
cp -rf README.md.template ../examples/README.md.template
echo -ne "[#---]"\\r
cp -rf README.md.template ../message/README.md.template
echo -ne "[##--]"\\r
cp -rf README.md.template ../network/README.md.template
echo -ne "[###-]"\\r
cp -rf README.md.template ../peer/README.md.template
echo -e "[####] done!"

# generate readme.md
echo "[TASK] Generating readme using autoreadme"

echo -ne "[-----]"\\r
$GOPATH/bin/autoreadme -f ../
echo -ne "[#----]"\\r
$GOPATH/bin/autoreadme -f ../examples
echo -ne "[##---]"\\r
$GOPATH/bin/autoreadme -f ../message
echo -ne "[###--]"\\r
$GOPATH/bin/autoreadme -f ../network
echo -ne "[####-]"\\r
$GOPATH/bin/autoreadme -f ../peer
echo -e "\r[#####] done!"

echo "[TASK] Rendering godoc to *.html"
echo -ne "[-----]"\\r
godoc -html -url "/pkg/splashp2p/" > ~/projects/apps/splashledgersite/public/pkg/splashp2p/index.html
echo -ne "[#----]"\\r
godoc -html -url "/pkg/splashp2p/examples/" > ~/projects/apps/splashledgersite/public/pkg/splashp2p/examples/index.html
echo -ne "[##---]"\\r
godoc -html -url "/pkg/splashp2p/message/" > ~/projects/apps/splashledgersite/public/pkg/splashp2p/message/index.html
echo -ne "[###--]"\\r
godoc -html -url "/pkg/splashp2p/network/" > ~/projects/apps/splashledgersite/public/pkg/splashp2p/network/index.html
echo -ne "[####-]"\\r
godoc -html -url "/pkg/splashp2p/peer/" > ~/projects/apps/splashledgersite/public/pkg/splashp2p/peer/index.html
echo "[#####] done!"


echo "[TASK] Commiting to .git"

if [ "$CommitMessage" != "" ]
then
git add .
git commit -m"$CommitMessage"
fi

echo "[DONE!]"