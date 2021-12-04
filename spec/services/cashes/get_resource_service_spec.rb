# frozen_string_literal: true

RSpec.describe Cashes::GetResourceService do
  subject { described_class.new params }

  let(:params) { { id: 27 } }

  describe '#cash' do
    before { allow(Cash).to receive(:find).with(27).and_return(:cash) }

    its(:cash) { should eq :cash }
  end

  describe '#call' do
    let(:cash) { stub_model Cash }

    before { allow(subject).to receive(:cash).and_return(cash) }

    its(:call) { should be_success }

    its('call.object') { should eq cash }
  end
end
