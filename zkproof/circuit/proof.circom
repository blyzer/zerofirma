pragma circom 2.0.0;

template SigProof() {
    signal input H;
    signal input C;
    signal input status;
    // simple example
    signal output out;
    out <== H + C + status;
}

component main = SigProof();
