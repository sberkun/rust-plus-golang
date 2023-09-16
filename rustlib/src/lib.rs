extern crate core;
extern crate libc;
extern crate ark_ec;
extern crate ark_bls12_377;
extern crate num_traits;

use core::slice;
use std::mem;

use ark_ec::AffineCurve;
use num_traits::identities::Zero;
use libc::c_void;
use ark_bls12_377::{Fr, G1Affine};


#[repr(C)]
pub struct GoG1Jac {
	x: [u64; 6],
    y: [u64; 6],
    z: [u64; 6]
}

#[no_mangle]
pub extern "C" fn multi_scalar_init_wrapper(points_ptr: *const c_void, len: u64) -> *mut c_void {
    let points: &[G1Affine] = unsafe {
        slice::from_raw_parts(points_ptr as *const G1Affine, len as usize)
    };
    println!("{:?}", points);
    println!();
    return points_ptr as *mut c_void;
}

#[no_mangle]
pub extern "C" fn multi_scalar_mult_wrapper(p: *mut c_void, ctx: *mut c_void, scalars: *const c_void, len: u64) {
    let scalars: &[Fr] = unsafe {
        slice::from_raw_parts(scalars as *const Fr, len as usize)
    };

    println!("banana {}", mem::size_of::<G1Affine>());
    println!("free banana {}", mem::size_of::<Fr>());
    let b: <G1Affine as AffineCurve>::Projective = 
        <G1Affine as AffineCurve>::Projective::zero();

    println!("{:?}", ctx);
    println!();
    println!("{:?}", scalars);
    println!();

    let return_value = GoG1Jac { x: [3,3,3,3,3,3], y: [1,1,1,1,1,1], z: [4,4,4,4,4,4] };
    unsafe {
        *(p as *mut GoG1Jac) = return_value;
    }
}

