# frozen_string_literal: true

RSpec.describe Tags::CreateService do
  subject { described_class.new category, params }

  let(:category) { build :category }

  let(:params) { acp(tag: { name: nil }) }

  its(:resource_params) { should eq params.require(:tag).permit! }

  describe '#tag' do
    context do
      before { subject.instance_variable_set :@tag, :tag }

      its(:tag) { should eq :tag }
    end

    context do
      before { allow(subject).to receive(:resource_params).and_return(:resource_params) }

      before { allow(category).to receive_message_chain(:tags, :new).with(:resource_params).and_return(:tag) }

      its(:tag) { should eq :tag }
    end
  end

  describe '#call' do
    let(:tag) { build :tag }

    before { allow(subject).to receive(:tag).and_return(tag) }

    context do
      before { allow(tag).to receive(:save).and_return(true) }

      its(:call) { should be_success }

      its('call.object') { should eq tag }
    end

    context do
      before { allow(tag).to receive(:save).and_return(false) }

      its(:call) { should be_failure }

      its('call.object') { should eq tag }
    end
  end
end
