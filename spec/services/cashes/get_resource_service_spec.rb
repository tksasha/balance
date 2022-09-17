# frozen_string_literal: true

RSpec.describe Cashes::GetResourceService do
  subject { described_class.new params }

  let(:params) { { id: 27 } }

  describe '#cash' do
    before { allow(Cash).to receive(:find).with(27).and_return(:cash) }

    its(:cash) { is_expected.to eq :cash }
  end

  describe '#call' do
    let(:cash) { stub_model Cash }

    before { allow(subject).to receive(:cash).and_return(cash) }

    its(:call) { is_expected.to be_success }

    its('call.object') { is_expected.to eq cash }
  end
end
