# frozen_string_literal: true

RSpec.describe Categories::CreateService do
  subject { described_class.new params }

  let(:params) { acp(category: { name: nil, income: nil, visible: nil, currency: nil }) }

  its(:resource_params) { should eq params.require(:category).permit! }

  describe '#category' do
    context do
      before { subject.instance_variable_set :@category, :category }

      its(:category) { should eq :category }
    end

    context do
      before { allow(subject).to receive(:resource_params).and_return(:resource_params) }

      before { allow(Category).to receive(:new).with(:resource_params).and_return(:category) }

      its(:category) { should eq :category }
    end
  end

  describe '#call' do
    let(:category) { stub_model Category }

    before { allow(subject).to receive(:category).and_return(category) }

    context do
      before { allow(category).to receive(:save).and_return(true) }

      its(:call) { should be_success }

      its('call.object') { should eq category }
    end

    context do
      before { allow(category).to receive(:save).and_return(false) }

      its(:call) { should be_failure }

      its('call.object') { should eq category }
    end
  end
end