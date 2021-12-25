# frozen_string_literal: true

RSpec.describe Websockets::UpdateAtEndService do
  subject { described_class.new currency }

  let(:currency) { 'uah' }

  describe '#at_end' do
    before { allow(CalculateAtEndService).to receive(:call).with('uah').and_return(241_513) }

    its(:at_end) { should eq '241 513.00' }
  end

  describe '#call' do
    before { allow(subject).to receive(:at_end).and_return('241 516.00') }

    before do
      expect(ActionCable)
        .to receive_message_chain(:server, :broadcast)
        .with('NotificationsChannel', { type: :at_end, value: '241 516.00' })
    end

    it { expect { subject.call }.not_to raise_error }
  end
end
