# frozen_string_literal: true

RSpec.describe Backoffice::Tags::UpdateService do
  subject { described_class.new params }

  let(:params) { acp(id: 12, tag: { name: nil }) }

  describe '#tag' do
    context do
      before { subject.instance_variable_set :@tag, :tag }

      its(:tag) { should eq :tag }
    end

    context do
      before { allow(Tag).to receive(:find).with(12).and_return(:tag) }

      its(:tag) { should eq :tag }
    end
  end

  its(:resource_params) { should eq params.require(:tag).permit! }

  describe '#call' do
    let(:tag) { build :tag }

    before { allow(subject).to receive(:tag).and_return(tag) }

    before { allow(subject).to receive(:resource_params).and_return(:resource_params) }

    context do
      before { allow(tag).to receive(:update).with(:resource_params).and_return(true) }

      its(:call) { should be_success }

      its('call.object') { should eq tag }
    end

    context do
      before { allow(tag).to receive(:update).with(:resource_params).and_return(false) }

      its(:call) { should be_failure }

      its('call.object') { should eq tag }
    end
  end
end
