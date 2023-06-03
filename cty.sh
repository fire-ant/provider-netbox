files=($(ls package/crds))

for f in ${files[@]}; do
    echo "$f"
    ./cty generate -c package/crds/$f -o examples/generated/
done

