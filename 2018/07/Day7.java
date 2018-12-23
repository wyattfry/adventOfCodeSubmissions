import java.util.ArrayList;
import java.util.HashMap;
import java.util.Map;
import java.util.Set;
import java.util.HashSet;
import java.util.List;

import adventOfCode.ReadTextFile;

public class Day7 {
  public static void main(String[] args) {
    String filename = "/Users/wyattfry/Projects/adventOfCode/07/input";
    run(filename);
  }

  public static void run(String filename) {
    String[] steps = ReadTextFile.toStringArray(filename);
    Map<String, Node> stepMap = new HashMap<>();
    String stepName;
    Node step;
    String nextStepName;
    Node nextStep;
    for (int i = 0; i < steps.length; i++) {
      stepName = steps[i].split(" ")[1];
      nextStepName = steps[i].split(" ")[7];
      // System.out.printf("%s <--> %s\n", stepName, nextStepName);
      if (!stepMap.containsKey(stepName)) {
        step = new Node(stepName);
        stepMap.put(stepName, step);
      } else {
        step = stepMap.get(stepName);
      }
      if (!stepMap.containsKey(nextStepName)) {
        nextStep = new Node(nextStepName);
        stepMap.put(nextStepName, nextStep);
      } else {
        nextStep = stepMap.get(nextStepName);
      }
      step.addNextNode(nextStep);
      nextStep.addPrevNode(step);
    }
    Node lastStep = null;
    for (Node n : stepMap.values()) {
      if (n.getNextNodes().size() == 0) {
        lastStep = n;
        break;
      }
    }


    List<Node> availableSteps = new ArrayList<>();
    String sequence = "";
    Node stepToDo;
    Map<String, Node> doneStepMap = new HashMap<>();
    while (stepMap.size() > 0) {
      for (Node n : stepMap.values()) {
        if (n.getPrevNodes().size() == 0) {
          if (!availableSteps.contains(n)) {
            availableSteps.add(n);
            // System.out.println("availableSteps " + availableSteps);
          }
        }
      }
      availableSteps.sort((step1, step2) -> step1.getName().compareTo(step2.getName()));
      stepToDo = availableSteps.get(0);
      sequence += stepToDo.getName();
      availableSteps.remove(stepToDo);
      for (Node n : stepToDo.getNextNodes()) {
        n.removePrevNode(stepToDo);
      }
      stepMap.remove(stepToDo.getName());
      doneStepMap.put(stepToDo.getName(), stepToDo);
    }
    System.out.println("Sequence: " + sequence);
    
    // find a node such that .preNodes.size() == 0
    // Node firstStep = null;
    // for (Node n : stepMap.values()) {
    //   if (n.getPrevNodes().size() == 0) {
    //     firstStep = n;
    //     break;
    //   }
    // }
    
    // List<Node> availableSteps = new ArrayList<>();
    // System.out.println("First step: " + firstStep.getName());
    // String sequence = firstStep.getName();
    // traverse(firstStep, availableSteps, sequence);
    // System.out.println("Sequence: " + sequence);

  } // run


  public static void traverse(Node startNode, List<Node> availableSteps, String sequence) {
    Node stepToRemove;
    if (startNode.getNextNodes().size() > 0) {
      availableSteps.addAll(startNode.getNextNodes());
      availableSteps.sort((s1, s2) -> s1.getName().compareTo(s2.getName()));
      stepToRemove = availableSteps.get(0);
      sequence += stepToRemove.getName();
      availableSteps.remove(stepToRemove);
      System.out.println(sequence + " " + availableSteps.toString());
      traverse(stepToRemove, availableSteps, sequence);
    }
  }
  
}

class Node {
  String name;
  Set<Node> nextNodes;
  Set<Node> prevNodes;
  public Node(String name) {
    this.name = name;
    this.nextNodes = new HashSet<>();
    this.prevNodes = new HashSet<>();
  }
  public String getName() {
    return this.name;
  }
  public void addNextNode(Node n) {
    this.nextNodes.add(n);
  }
  public void addPrevNode(Node n) {
    this.prevNodes.add(n);
  }
  public Set<Node> getNextNodes() {
    return this.nextNodes;
  }
  public Set<Node> getPrevNodes() {
    return this.prevNodes;
  }
  public boolean removeNextNode(Node n) {
    return this.nextNodes.remove(n);
  }
  public boolean removePrevNode(Node n) {
    return this.prevNodes.remove(n);
  }
  @Override
  public String toString() {
    return String.format("[Node %s]", name);
  }
}