FROM alpine

ADD chalkboard ./chalkboard
ADD static ./static

ENTRYPOINT ["./chalkboard"]