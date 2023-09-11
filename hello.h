// NOTE: You could use https://michael-f-bryan.github.io/rust-ffi-guide/cbindgen.html to generate
// this header automatically from your Rust code.  But for now, we'll just write it by hand.


struct Dummy {
    char* thing1;
    char* thing2;
};


void hello(char *name);
void whisper(char *message);
void do_thing(struct Dummy* d);
