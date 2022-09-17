# frozen_string_literal: true

RSpec.describe Categories::GetResourceService do
  subject { described_class.new params }

  let(:params) { { id: 9 } }

  describe '#category' do
    context do
      before { subject.instance_variable_set :@category, :category }

      its(:category) { is_expected.to eq :category }
    end

    context do
      before { allow(Category).to receive(:find).with(9).and_return(:category) }

      its(:category) { is_expected.to eq :category }
    end
  end

  describe '#call' do
    let(:category) { stub_model Category }

    before { allow(subject).to receive(:category).and_return(category) }

    its(:call) { is_expected.to be_success }

    its('call.object') { is_expected.to eq category }
  end
end
