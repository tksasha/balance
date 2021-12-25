# frozen_string_literal: true

RSpec.describe Websockets::UpdateBalanceService do
  subject { described_class.new currency }

  let(:currency) { 'uah' }

  describe '#balance' do
    before { allow(CalculateBalanceService).to receive(:call).with('uah').and_return(241_513) }

    its(:balance) { should eq '241 513.00' }
  end

  describe '#call' do
    before { allow(subject).to receive(:balance).and_return('241 516.00') }

    before do
      expect(ActionCable)
        .to receive_message_chain(:server, :broadcast)
        .with('NotificationsChannel', { type: :balance, value: '241 516.00' })
    end

    it { expect { subject.call }.not_to raise_error }
  end
end
