const go = new Go(); // Go 런타임 초기화

async function loadWasm() {
  const wasmModule = await WebAssembly.instantiateStreaming(
    fetch("js/main.wasm"),
    go.importObject
  );
  go.run(wasmModule.instance);
}

function getValues() {
  const value1 = parseFloat(document.getElementById("value1").value);
  const value2 = parseFloat(document.getElementById("value2").value);
  const output = document.getElementById("output");

  if (isNaN(value1) || isNaN(value2)) {
    output.value = "숫자를 입력하세요.";
    return null;
  }

  return { value1, value2, output };
}

function performAddition() {
  const values = getValues();
  if (!values) return;

  const result = add(values.value1, values.value2);
  values.output.value = `덧셈 결과: ${result}`;
}

function performSubtraction() {
  const values = getValues();
  if (!values) return;

  const result = subtract(values.value1, values.value2);
  values.output.value = `뺄셈 결과: ${result}`;
}

function performMultiplication() {
  const values = getValues();
  if (!values) return;

  const result = multiply(values.value1, values.value2);
  values.output.value = `곱셈 결과: ${result}`;
}

function performDivision() {
  const values = getValues();
  if (!values) return;

  if (values.value2 === 0) {
    values.output.value = "0으로 나눌 수 없습니다.";
    return;
  }

  const result = divide(values.value1, values.value2);
  values.output.value = `나눗셈 결과: ${result}`;
}

loadWasm();
