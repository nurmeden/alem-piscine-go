INTT=$(head -n 179 streets/Buckingham_Place | tail -n 1 | cut -d '#' -f2)
echo "$INTT"
cat ./interviews/interview-"$INTT"
echo "$MAIN_SUSPECT"