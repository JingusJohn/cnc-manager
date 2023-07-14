export const load = async () => {
  console.log("getting data?")

  const getMessage = async () => {
    const res = await fetch("http://gateway:1323");
    console.log("result: ", res);
    if (res.status == 200) {
      const item = await res.text()
      return item;

    } else {
      return res;
    }
  }

  return {
    message: await getMessage()
  }
}
