# frozen_string_literal: true

RSpec.describe ItemSearcher do
  describe '#params' do
    subject { described_class.new relation, params }

    let(:relation) { double }

    let(:params) { { currency: 'usd' } }

    let(:month) { Month.now }

    before { allow(subject).to receive(:month).and_return(month) }

    its(:params) { is_expected.to eq currency: 'usd', month: }
  end
end