# frozen_string_literal: true

RSpec.describe ItemSearcher do
  describe '#params' do
    let(:relation) { double }

    let(:params) { { currency: 'usd' } }

    let(:month) { Month.today }

    before { allow(subject).to receive(:month).and_return(month) }

    subject { described_class.new relation, params }

    its(:params) { should eq currency: 'usd', month: month }
  end
end
