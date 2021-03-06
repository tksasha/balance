# frozen_string_literal: true

RSpec.describe ParseMonthService do
  let(:params) { { month: '2021-03' } }

  subject { described_class.new params }

  describe '#month' do
    its(:month) { should eq Month.new(2021, 3) }

    context do
      let(:params) { {} }

      before { travel_to Date.new(2021, 4, 1) }

      after { travel_back }

      its(:month) { should eq Month.new(2021, 4) }
    end
  end

  describe '#call' do
    let(:month) { Month.today }

    before { allow(subject).to receive(:month).and_return(month) }

    its(:call) { should eq month }
  end

  describe '.call' do
    subject { described_class.call params }

    before do
      #
      # described_class.new(params).call
      #
      allow(described_class).to receive(:new).with(params) do
        double.tap do |a|
          allow(a).to receive(:call)
        end
      end
    end

    it { expect { subject }.not_to raise_error }
  end
end
