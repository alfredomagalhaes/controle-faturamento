import { Modal, DatePicker, InputNumber, Form} from 'antd';


export function ModalEdicaoFaturamento(props){
    const formatoAnoMes = 'MM/YYYY'
    return (
        <Modal 
            title="Faturamento" 
            visible={props.isOpen} 
            onOk={props.onOkFunc} 
            onCancel={props.onCancelFunc}
        >
            <Form
                labelCol={{ span: 6 }}
                wrapperCol={{ span: 14 }}
                layout="horizontal"
            >
                <Form.Item label="Selecione período">
                    <DatePicker 
                        format={formatoAnoMes} 
                        picker="month"
                        placeholder ="MM/AAAA"
                        onChange={props.onChangeAnoMes}
                    />
                </Form.Item>
                <Form.Item label="Valor Faturado">
                    <InputNumber 
                        defaultValue={0}
                        style={{ width: 156 }}
                        value={props.valorFat}
                        onChange={props.onChangeValorFat}
                        decimalSeparator=","
                    />
                </Form.Item>

            </Form> 
            
        </Modal>
    );
}