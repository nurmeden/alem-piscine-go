INTT=$(head -n 179 streets/buckingham_Place | tail -n 1 | cut -d '#' -f2)
echo "$INTT"
cat ./interviews/interview-"$INTT"
echo "$NAIN_SUSPECT"