extern crate core;
extern crate libc;

use core::slice;

use libc::c_void;

type Element = [u64; 6];

#[repr(C)]
#[derive(Debug)]
pub struct G1Affine {
    x: Element,
    y: Element
}

#[repr(C)]
pub struct G1Jac {
	x: Element,
    y: Element,
    z: Element
}

#[no_mangle]
pub extern "C" fn multi_scalar_mult_wrapper(p: *mut c_void, ctx: *mut c_void, points: *const c_void, scalars: *const c_void, len: u64) {
    let points: &[G1Affine] = unsafe {
        slice::from_raw_parts(points as *const G1Affine, len as usize)
    };
    let scalars: &[Element] = unsafe {
        slice::from_raw_parts(scalars as *const Element, len as usize)
    };

    println!("{:?}", ctx);
    println!("{:?}", points);
    println!("{:?}", scalars);

    let return_value = G1Jac { x: [3,3,3,3,3,3], y: [1,1,1,1,1,1], z: [4,4,4,4,4,4] };
    unsafe {
        *(p as *mut G1Jac) = return_value;
    }
}

#[no_mangle]
pub extern "C" fn multi_scalar_init_wrapper(points: *const c_void, len: u64) -> *mut c_void {
    return 0x12345 as *mut c_void;
}
